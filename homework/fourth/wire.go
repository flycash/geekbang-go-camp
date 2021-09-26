//+build wireinject

package main

import (
	"geekbang/geekbang-go-camp/homework/fourth/internal/biz"
	"geekbang/geekbang-go-camp/homework/fourth/internal/data"
	"geekbang/geekbang-go-camp/homework/fourth/internal/service"
	"github.com/google/wire"
)

func InitUserService() *service.UserService {
	wire.Build(service.NewUserService, biz.NewUserBiz, data.NewUserRepo, data.NewDB, data.NewCache)
	return &service.UserService{}
}