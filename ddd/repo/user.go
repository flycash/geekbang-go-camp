package repo

import (
	"geekbang/geekbang-go-camp/ddd/domain/user"
)

func SaveUser(user user.UserAggregate) {
	// 要分成两步，首先保存 UserEntity - DAO
	// 其次保存 Address
	// 这里就会依赖于 DAO 和 Cache
}

type UserRepo interface {
	SaveUser(user user.UserAggregate)
}

type UserRepoDao struct {
}

func (u UserRepoDao) SaveUser(user user.UserAggregate) {
	// TODO implement me
	panic("implement me")
}

type UserRepoWithCache struct {
	UserRepo
}

func (u UserRepoWithCache) SaveUser(user user.UserAggregate) {
	u.UserRepo.SaveUser(user)
	// 一堆缓存操作
	panic("implement me")
}
