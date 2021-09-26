package data

import (
	"github.com/gotomicro/ego-component/egorm"
	"github.com/gotomicro/ego-component/eredis"
)

type UserRepo struct {
	db *egorm.Component
	cache *eredis.Component
}

func NewUserRepo(db *egorm.Component, cache *eredis.Component) *UserRepo {
	return &UserRepo{
		db: db,
		cache: cache,
	}
}

func (u *UserRepo) GetUser(id uint64) (*UserPO, error) {
	return &UserPO{
		Nickname: "Tom",
	}, nil
}

type UserPO struct {
	Nickname string
}


