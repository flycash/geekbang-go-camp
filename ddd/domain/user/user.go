package user

type UserEntity struct {
	Id int64
}

type Address struct {
}

type UserAggregate struct {
	// 聚合根我们一般用组合的形式
	UserEntity
	Addr []Address
}

// func (u UserAggregate) Validate() error {
// 	u.UserEntity.Validate()
// 	for _, a := range u.Addr {
// 		a.Validate()
// 	}
// 一个个字段检查过去
// 一个个值对象检查过去
// 	panic("implement me")
// }

// func (u OrderAggregate) Validate() error {
// 	// 一个个字段检查过去
// 	// 一个个值对象检查过去
// 	u.TotalAmount.Validate()
// 	panic("implement me")
// }
//
// type Amount struct {
// 	amt      uint64
// 	currency string
// }
//
// func (u Amount) Validate() error {
// 	// 一个个字段检查过去
// 	// 一个个值对象检查过去
// 	panic("implement me")
// }
