package event

import (
	"geekbang/geekbang-go-camp/ddd/domain/user"
)

type UserRegisterEvent struct {
}

func PublishRegisterEvent(user user.UserAggregate) {

}
