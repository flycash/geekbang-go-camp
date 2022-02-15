package eleventh

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/pingcap/parser"
	"github.com/pingcap/parser/ast"
	"github.com/pingcap/parser/format"
	"github.com/pingcap/parser/model"
	_ "github.com/pingcap/tidb/types/parser_driver"
	"go.uber.org/zap/buffer"
)

// Sharding 整个 sharding 的过程
// 其实就是接收一个 sql 和 参数，然后我们执行它，返回结果
// 我们也可以不用自己的接口，而是实现一个 sharding 的 driver
// 类似于 MySQL 的 driver
// 出于教学的目的，这里我们尝试自己定义接口，会更加清晰
type Sharding interface {
	Exec(ctx context.Context, sql string, args ...interface{}) (sql.Result, error)
	Query(ctx context.Context, sql string, args ...interface{}) (*sql.Rows, error)
}

type mockShard struct {
	rewriter Rewriter
	executor Executor
	merger   Merger
}

func newMockShard() *mockShard {
	return &mockShard{
		rewriter: newAstRewriter(),
		executor: &simpleExecutor{},
		merger:   &dispatcherMerger{},
	}
}

func (m *mockShard) Query(ctx context.Context, sql string, args ...interface{}) (*sql.Rows, error) {
	rwRes, err := m.rewriter.Rewrite(ctx, &RewriteContext{
		sql:  sql,
		args: args,
	})
	if err != nil {
		return nil, err
	}
	exeRes, err := m.executor.Execute(ctx, &ExecuteContext{queries: rwRes.queries})
	if err != nil {
		return nil, err
	}

	mergeRes, err := m.merger.Merge(ctx, &MergeContext{
		results: exeRes.results,
	})
	if err != nil {
		return nil, err
	}
	return mergeRes.rows, err
}

// Exec 因为我们说 sharding 核心就是三个步骤
// 重写、执行和合并结果
// 于是我们引入三个接口来代表这三个过程
func (m *mockShard) Exec(ctx context.Context, sql string, args ...interface{}) (sql.Result, error) {
	rwRes, err := m.rewriter.Rewrite(ctx, &RewriteContext{
		sql:  sql,
		args: args,
	})
	if err != nil {
		return nil, err
	}
	exeRes, err := m.executor.Execute(ctx, &ExecuteContext{queries: rwRes.queries})
	if err != nil {
		return nil, err
	}

	mergeRes, err := m.merger.Merge(ctx, &MergeContext{
		results: exeRes.results,
	})
	if err != nil {
		return nil, err
	}
	return mergeRes.result, err
}

// Rewriter 代表重写 SQL
type Rewriter interface {
	Rewrite(ctx context.Context, rwCtx *RewriteContext) (*RewriteResult, error)
}

// RewriteContext 代表一个重写上下文，
// 里面放着你需要的各种数据。
type RewriteContext struct {
	sql  string
	args []interface{}
}

// RewriteResult 代表重写后的结果
type RewriteResult struct {
	// 重写后的一堆查询
	queries []*RewriteQuery
}

// RewriteQuery 重写后的 SQL
type RewriteQuery struct {
	sql  string
	args []interface{}

	// 你可能需要一些查询特征字段，用于执行和合并结果阶段使用
	// 比如说要根据 dbname 去找到链接信息，
	// 特别是要考虑到后面合并结果的方式五花八门，
	// 这里的字段可能会非常丰富
	tableName string
	dbName    string
}

// astRewriter 比如说利用 AST 来实现重写
type astRewriter struct {
	cfg *ShardingConfig
	p   *parser.Parser
}

type tableNameRewriterVisitor struct {
	tblName   *ast.TableName
	args      []interface{}
	shardFunc func(i interface{}) int
}

func (a *tableNameRewriterVisitor) Enter(n ast.Node) (node ast.Node, skipChildren bool) {
	if nd, ok := n.(*ast.TableName); ok {
		sharded := a.shardFunc(a.args[0])
		original := nd.Name.String()
		nd.Name = model.NewCIStr(fmt.Sprintf("%s_%d", original, sharded))
		a.tblName = nd
	}
	// 其实在这里，我们还需要解析 WHERE 部分，确认里面有没有 sharding key，
	// 并且要确定 sharding key 对应的查询条件
	// 以及查询条件对应的参数
	// 在这个例子里面，就是找到 id=? 中问号对应的参数
	// 但是因为我们只是例子，这里就不找了
	return n, false
}

func (a *tableNameRewriterVisitor) Leave(n ast.Node) (node ast.Node, ok bool) {
	return n, true
}

func newAstRewriter() *astRewriter {
	return &astRewriter{
		p: parser.New(),
	}
}

func (a *astRewriter) Rewrite(ctx context.Context, rwCtx *RewriteContext) (*RewriteResult, error) {
	// 在这里，构建起 AST 树
	// 修改 AST 的节点。比如说插入主键节点，然后节点的值就是主键生成策略生成的值
	// 最后将 AST 转为一个 SQL

	stmtNodes, _, err := a.p.Parse(rwCtx.sql, "", "")
	if err != nil {
		return nil, err
	}

	// 代表的是直接的根节点
	root := stmtNodes[0]
	visitor := &tableNameRewriterVisitor{
		args: rwCtx.args,
		// 这种如何 shard 的问题，应该是从配置里面解析生成，这里我们直接写死先
		shardFunc: func(i interface{}) int {
			return i.(int) % 2
		},
	}
	root.Accept(visitor)
	bytes := &buffer.Buffer{}
	_ = root.(*ast.SelectStmt).Restore(format.NewRestoreCtx(format.DefaultRestoreFlags, bytes))
	s := bytes.String()
	fmt.Printf("after rewrite: %s \n", s)
	return &RewriteResult{
		queries: []*RewriteQuery{{
			sql:       s,
			args:      rwCtx.args,
			tableName: visitor.tblName.Name.String(),
			dbName:    visitor.tblName.Schema.String(),
		}},
	}, nil
}

type ShardingConfig struct {
	// 这里就是你的各种 sharding 的配置
	// 比如说你的表怎么sharding
	// db 怎么 sharding
	// 主键生成策略是什么
	// 不同db的连接信息
	// ... 可以参考 shardingsphere 的配置文件，非常丰富
}

// Executor 代表执行 SQL
type Executor interface {
	// Execute 这里额外返回一个错误，是我们自身的错误，而不是执行查询引起的错误
	Execute(ctx context.Context, exeCtx *ExecuteContext) (*ExecuteResult, error)
}

// ExecuteContext 代表一个执行上下文
type ExecuteContext struct {
	queries []*RewriteQuery
}

// ExecuteResult 执行结果
type ExecuteResult struct {
	results []*QueryResult
}

type QueryResult struct {
	// 合并结果的时候，Merger 的实现自己知道该用什么字段不该用什么字段
	queryRows *sql.Rows
	err       error
	execRes   sql.Result

	// 可以考虑改进接口，也可以直接在这里保留
	query *RewriteQuery
}

// 简单的遍历执行
type simpleExecutor struct {
	// 维持住了所有的物理表创建的 DB
	// 它基本上是在初始化的时候根据配置来创建的
	dbConn map[string]*sql.DB
}

func (p *simpleExecutor) Execute(ctx context.Context, exeCtx *ExecuteContext) (*ExecuteResult, error) {
	queryResult := make([]*QueryResult, 0, len(exeCtx.queries))
	for _, query := range exeCtx.queries {
		db, ok := p.dbConn[query.dbName]
		if !ok {
			// 要么是用户没有配置，要么是 sharding 出错了
			return nil, errors.New("invalid sharding queries")
		}

		// 这里要判断是SELECT 还是 UPDATE 之类的
		res, err := db.ExecContext(ctx, query.sql, query.args...)
		// 或者是
		// rows, err := db.QueryContext(ctx, query.sql, query.args...)
		queryResult = append(queryResult, &QueryResult{query: query, err: err, execRes: res})
	}
	return &ExecuteResult{results: queryResult}, nil
}

// Merger 代表合并结果
// 这个接口会有很多很多的实现，
type Merger interface {
	Merge(ctx context.Context, mergeCtx *MergeContext) (*MergeResult, error)
}

type MergeContext struct {
	results []*QueryResult
}

type MergeResult struct {
	rows   *sql.Rows
	result sql.Result
	error  error
}

// 门面模式
type dispatcherMerger struct {
	// 处理平均值的
	avgMerger Merger

	// 处理计数的
	cntMerger Merger

	// ... 你会有一大堆

}

func (d *dispatcherMerger) Merge(ctx context.Context, mergeCtx *MergeContext) (*MergeResult, error) {
	// 检查 mergeCtx 里面的每一个查询结果和查询特征，然后选择
	// if queryCnt(mergeCtx) {
	//     return cntMerger.Merge(ctx, mergeCtx)
	// }
	panic("implement me")
}
