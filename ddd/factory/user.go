package factory

import (
	"geekbang/geekbang-go-camp/ddd/domain/user"
	"net/http"
)

func CreateNewUser(r http.Request) user.UserAggregate {
	panic("implement me")
}
