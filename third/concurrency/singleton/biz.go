package singleton

import "geekbang/geekbang-go-camp/third/concurrency"

var aa = "aa"

func SetAA(a string) {
	aa = a
}

var defaultHandler MyHandler

var s = concurrency.GetSingleInstance()

func init() {
	a := s
	a.Single()
}

func Handle() {
	s.Single()
	// defaultHandler.Handle()
}

func InitDefault(config string) {
	// defaultHandler = xxxx
}

type MyHandler interface {
	Handle()
}
