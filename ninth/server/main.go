package main

import (
	"geekbang/geekbang-go-camp/ninth/rpc"
)

func main() {
	tracingBuilder := &rpc.ServerSideTracingFilterBuilder{}

	svr := rpc.NewToyProtocolServer("127.0.0.1:8081",
		rpc.ExtractContextFilterBuilder, tracingBuilder.Build)
	svr.RegisterService(&UserService{})
	if err := svr.Start(); err != nil {
		panic(err)
	}
}
