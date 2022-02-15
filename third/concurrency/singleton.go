package concurrency

import (
	"fmt"
	"sync"
)

// 可以定义一个接口
// type Single interface {
// 	Single()
// }

type singleton struct {
	// 有很多字段
}

func (s *singleton) Single() {
	fmt.Println("I am single")
}

var instance *singleton
var instanceOnce sync.Once

// GetSingleInstance 返回接口
func GetSingleInstance() *singleton {
	instanceOnce.Do(func() {
		instance = &singleton{}
	})
	return instance
}

type MyBiz struct {
	once sync.Once
}

// 只被执行一次
func (m *MyBiz) Init() {
	m.once.Do(func() {

	})
}
