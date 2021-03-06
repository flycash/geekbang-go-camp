// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"geekbang/geekbang-go-camp/homework/fourth/internal/biz"
	"geekbang/geekbang-go-camp/homework/fourth/internal/data"
	"geekbang/geekbang-go-camp/homework/fourth/internal/service"
)

// Injectors from wire.go:

func InitUserService() *service.UserService {
	db := data.NewDB()
	component := data.NewCache()
	userRepo := data.NewUserRepo(db, component)
	userBiz := biz.NewUserBiz(userRepo)
	userService := service.NewUserService(userBiz)
	return userService
}
