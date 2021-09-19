package singleton

var aa = "aa"

func SetAA(a string) {
	aa = a
}
func Singleton()  {
	// ...
}

var defaultHandler MyHandler

func init() {
	defaultHandler= new()
}

func Handle() {
	defaultHandler.Handle()
}

func InitDefault(config string)  {
	defaultHandler = xxxx
}

type MyHandler interface {
	Handle()
}

