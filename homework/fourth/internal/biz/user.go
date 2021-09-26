package biz

import (
	"errors"
	"geekbang/geekbang-go-camp/homework/fourth/internal/data"
)

type UserBiz struct {
	repo *data.UserRepo
}

func NewUserBiz(repo *data.UserRepo) *UserBiz {
	return &UserBiz{
		repo: repo,
	}
}

func(ub *UserBiz) GetUserById(uid uint64) (*UserDO, error){
	if uid == 0 {
		// 这里，你可能要考虑使用错误码，是内部业务错误码
		return nil, errors.New("invalid user id")
	}
	u, err := ub.repo.GetUser(uid)
	if err != nil {
		// 理论上来说，repo 会把 error 组装好，附加上各种必要的debug 信息，这里可以直接返回
		// 如果 repo 里面并没有处理，依旧是保留着原生的DB错误数据，这边要考虑转换具体业务错误
		// 比如说 NoRows 这种错误，要考虑转换为 user not found 或者 invalid user id
		return nil, err
	}
	return &UserDO{
		Nickname: u.Nickname,
	}, nil
}

type UserDO struct {
	Nickname string
}

type MyService interface {

}

type Option func(db *service) MyService

func NewDB(opts...Option) MyService {
	return &service{}
}

type service struct {

}