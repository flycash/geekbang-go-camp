package service

import (
	"geekbang/geekbang-go-camp/ddd/domain/user"
	"geekbang/geekbang-go-camp/ddd/event"
	"geekbang/geekbang-go-camp/ddd/factory"
	"geekbang/geekbang-go-camp/ddd/repo"
	"net/http"
)

func Register(req http.Request) user.UserAggregate {
	user := factory.CreateNewUser(req)
	user.Validate()
	// 过程式
	// Validate(user)
	repo.SaveUser(user)
	event.PublishRegisterEvent(user)
	return user
}
