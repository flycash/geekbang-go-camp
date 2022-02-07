package rpc

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"geekbang/geekbang-go-camp/ninth/dto"
	"net"
	"reflect"
)

type Server interface {
	Start() error
	// RegisterService requires registering all your services before start
	RegisterService(s Service)
}
type ProviderProxyFactory func() ProviderProxy

type toyProtocolServer struct {
	address  string
	services map[string]*reflectionStub
	invoker  Proxy
}

func (t *toyProtocolServer) Start() error {
	ln, err := net.Listen("tcp", t.address)
	if err != nil {
		return err
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Printf("accept connection got error: %v", err)
		}
		go t.handleConnection(conn)
	}
}

func (t *toyProtocolServer) handleConnection(conn net.Conn) {
	for {
		// _ = conn.SetReadDeadline(time.Now().Add(10 * time.Second))
		bytes, err := ReadMsg(conn)
		if err != nil {
			return
		}
		// go func() {
		u := &dto.Request{}
		err = json.Unmarshal(bytes, u)
		resp, er := t.invoker.Invoke(context.Background(), u)
		// resp, er := t.Invoke(context.Background(), u)
		if resp == nil {
			resp = &dto.Response{}
		}
		if er != nil && len(resp.Error) == 0 {
			resp.Error = er.Error()
		}
		resp.MessageId = u.MessageId
		encode, er := EncodeProto(resp)
		if er != nil {
			fmt.Printf("encode resp failed: %v", er)
			return
		}
		_, er = conn.Write(encode)
		if er != nil {
			fmt.Printf("sending response failed: %v", er)
		}
		// }()
	}
}

func (t *toyProtocolServer) Invoke(ctx context.Context, req *dto.Request) (*dto.Response, error) {
	resp := &dto.Response{
		MessageId: req.MessageId,
	}
	s, ok := t.services[req.ServiceName]
	if !ok {
		return resp, errors.New("service not found")
	}
	respData, err := s.invoke(ctx, req.Method, req.Data)
	if err != nil {
		return resp, err
	}
	resp.Data = respData
	return resp, nil
}

func (t *toyProtocolServer) RegisterService(s Service) {
	t.services[s.ServiceName()] = &reflectionStub{
		s:     s,
		value: reflect.ValueOf(s),
	}
}

func NewToyProtocolServer(address string, fts ...FilterChain) Server {
	res := &toyProtocolServer{
		address:  address,
		services: make(map[string]*reflectionStub, 4),
	}
	ftl := NewFilterProxy(res, fts...)
	res.invoker = ftl
	return res
}

func ExtractContextFilterBuilder(next Filter) Filter {
	return func(ctx context.Context, req *dto.Request) (*dto.Response, error) {
		fmt.Printf("this is meta extract filter\n")
		for k, v := range req.Meta {
			ctx = context.WithValue(ctx, k, v)
		}
		return next(ctx, req)
	}
}
