package dto

type Request struct {

	// for scalability
	ServiceName string
	Method      string
	// request itself
	Data      []byte
	MessageId uint64
	// 这个 meta 就是用来传递链路元数据的，比如说课程上说的多租户或者 a/b test，链路超时之类的
	Meta map[string]string
}

type Response struct {
	Data      []byte
	Error     string
	MessageId uint64
	Meta      map[string]string
}

type FindByUserIdReq struct {
	Id uint64
}

type FindByUserIdResp struct {
	User *User
}

type User struct {
	Id         uint64
	Name       string
	Avatar     string
	Email      string
	Password   string
	CreateTime int
}
