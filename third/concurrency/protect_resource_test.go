package concurrency

import "testing"

func TestRegistry_Get(t *testing.T) {


	registry := &Registry{
		resources: map[string]interface{}{},
	}

	// 要求在应用启动前全部注册好
	registry.Register("a", "a-r")
	registry.Register("b", "b-r")
	registry.Register("c", "c-r")
	registry.Register("d", "d-r")
	registry.Register("e", "e-r")
	runApp()

	// ORM 框架 - 注册模型
	// Web 框架 - 注册路由

	// 前面只写，串行写，无并发写
	// ----------------------- 分界线 ---------------
	// 后面只读，可以并发读
}

func runApp()  {
	// 这里对资源的并发操作
}