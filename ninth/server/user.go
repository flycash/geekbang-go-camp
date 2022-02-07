package main

import (
	"context"
	"geekbang/geekbang-go-camp/ninth/dto"
	"time"
)

type UserService struct {
}

func (u *UserService) FindById(ctx context.Context, req *dto.FindByUserIdReq) (*dto.FindByUserIdResp, error) {

	return &dto.FindByUserIdResp{
		User: &dto.User{
			Id:         12,
			Name:       "Tom",
			Avatar:     "http://my-avatar",
			Email:      "xxx@xxx.com",
			Password:   "123456",
			CreateTime: time.Now().Second(),
		},
	}, nil
}

func (u *UserService) ServiceName() string {
	return "user"
}
