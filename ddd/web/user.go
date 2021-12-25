package web

import (
	"fmt"
	"geekbang/geekbang-go-camp/ddd/service"
	"net/http"
)

func Register(req http.Request) {
	user := service.Register(req)
	// 准备前端所需要的各种字段
	fmt.Printf("success %d", user.Id)

	// 有些人
	// user := factory.CreateNewUser(req)
	// user.Validate()
}
