package main

import (
	"context"
	"encoding/json"
	"fmt"
	"geekbang/geekbang-go-camp/ninth/dto"
	"geekbang/geekbang-go-camp/ninth/rpc"
)

func main() {
	us := &UserService{}
	tracingBuilder := &rpc.ClientSideTracingFilterBuilder{}
	metaInjectFilter := &rpc.InjectContextFilterBuilder{
		Reader: tracingBuilder.CtxReader,
	}
	err := rpc.InitToyProtocolProxy("0.0.0.0:8081",
		us, tracingBuilder.Build, metaInjectFilter.Build)
	if err != nil {
		panic(err)
	}
	resp, err := us.FindById(context.Background(), &dto.FindByUserIdReq{
		Id: 12,
	})
	if err != nil {
		panic(err)
	}
	data, _ := json.Marshal(resp)
	fmt.Printf("the response is: %s", data)
}

// 声明服务，反射会把 FindById 转成一个 RPC 调用
type UserService struct {
	FindById func(ctx context.Context, req *dto.FindByUserIdReq) (*dto.FindByUserIdResp, error)
}

func (u *UserService) ServiceName() string {
	return "user"
}
