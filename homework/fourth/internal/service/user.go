package service

import (
	"context"
	"geekbang/geekbang-go-camp/homework/fourth/api"
	"geekbang/geekbang-go-camp/homework/fourth/internal/biz"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UserService struct {
	api.UnimplementedUserServiceServer
	biz *biz.UserBiz
}

func NewUserService(biz *biz.UserBiz) *UserService {
	return &UserService{
		biz: biz,
	}
}

func (u *UserService) UserInfo(ctx context.Context, request *api.UserInfoRequest) (*api.UserInfoReply, error) {
	usr, err := u.biz.GetUserById(request.Uid)
	if err != nil {
		// 这里要考虑尝试转化为携带了 error code 的 error
		// 这里要注意，如果在 biz 的地方定义的是内部业务错误码，即你可能不希望外界知道你的内部业务错误码，这里就要转化
		// 如果你的错误码是不介意暴露出去的，那么可以不用转换，直接 return。
		// 另外，如果你不希望下层依赖于 grpc，而这里又要利用 grpc 的错误码传递，那么还是需要做转化
		return nil, status.Error(codes.Code(1), "system error")
	}

	return &api.UserInfoReply{
		User: &api.User{
			Nickname: usr.Nickname,
		},
	}, nil
}
