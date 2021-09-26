package main

import (
	"geekbang/geekbang-go-camp/homework/fourth/api"
	"geekbang/geekbang-go-camp/homework/fourth/internal/task"
	"github.com/gotomicro/ego"
	"github.com/gotomicro/ego/core/elog"
	"github.com/gotomicro/ego/server/egrpc"
	"github.com/gotomicro/ego/task/ejob"
)
//  export EGO_DEBUG=true && go run main.go --config=configs/local.toml
func main() {

	// ego 自身完成了监听信号，然后退出的功能
	if err := ego.New().Serve(func() *egrpc.Component {
		srv := egrpc.Load("server.grpc").Build()
		api.RegisterUserServiceServer(srv, InitUserService())
		return srv
	}()).Job(ejob.Job("say_hello", task.Hello)).Run(); err != nil {
		elog.Panic("startup", elog.FieldErr(err))
	}
}

