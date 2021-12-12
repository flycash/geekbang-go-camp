package second

import (
	"database/sql"
	"fmt"
	"github.com/pkg/errors"
	"strings"
)

// 用来保证和第三方库解耦，这里的第三方库是 sql 包
var NotFound = errors.New("not found")
var notFoundCode = 40001
var systemErr = 50001

func Biz() error {
	err := Dao("")
	// 只能依赖于 DAO 的错误，也就是这里不能检测 sql.ErrXXX
	if errors.Is(err, NotFound) {
		// 没找到，这时候要站在业务的角度考虑，如果找不到是正常的，
		// 比如说一个新用户没有订单，这是正常的，那么应该 return nil。同时考虑加一个 debug 日志
		// 如果说业务上认为这必然是有的，比如说用户点击订单详情，但是这里居然没找到
		// 那就要考虑转为另外一个业务错误了，或者直接返回错误响应
		return nil
	}
	if err != nil {
		// 不管怎么说，出现了数据库查询的问题，可以转为业务领域错误，也可以继续向上传递
	}
	return nil
}

func Dao(query string) error {
	err := mockError()

	if err == sql.ErrNoRows {
		// 在这一步封装好查询参数，这样DEBUG就能知道请求什么数据，没找到
		// 同时带上了堆栈信息方便定位
		return errors.Wrapf(NotFound, fmt.Sprintf("data not found, sql: %s ", query))
	}

	if err != nil {
		return errors.Wrapf(err, fmt.Sprintf("db query system error sql: %s ", query))
	}
	// do something
	return nil
}

func Biz1() error {
	err := Dao("")

	if err != nil {
		// 不管怎么说，出现了数据库查询的问题，可以转为业务领域错误，也可以继续向上传递
	}
	return nil
}

func Dao1(query string) error {
	err := mockError()

	if err != nil {
		// 在这一步封装好查询参数，这样DEBUG就能知道请求什么数据，没找到
		// 同时带上了堆栈信息方便定位
		// 我们没有仔细区别 err 是什么，反正就是告诉上游，出错了
		return errors.Wrapf(NotFound, fmt.Sprintf("data not found, sql: %s ", query))
	}
	// do something
	return nil
}

func Biz2() error {
	err := Dao("")

	if IsNoRow(err) {
		// 不管怎么说，出现了数据库查询的问题，可以转为业务领域错误，也可以继续向上传递
	} else if err != nil {

	}
	return nil
}

func Dao2(query string) error {
	err := mockError()

	if err == sql.ErrNoRows {
		// 在这一步封装好查询参数，这样DEBUG就能知道请求什么数据，没找到
		// 同时带上了堆栈信息方便定位
		// 我们没有仔细区别 err 是什么，反正就是告诉上游，出错了
		return fmt.Errorf("%d, not found", notFoundCode)
	} else if err != nil {
		return fmt.Errorf("%d, not found", systemErr)
	}
	// do something
	return nil
}

func IsNoRow(err error) bool {
	return strings.HasPrefix(err.Error(), fmt.Sprintf("%d", notFoundCode))
}

func mockError() error {
	return sql.ErrNoRows
}
