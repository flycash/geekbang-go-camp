package rpc

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"geekbang/geekbang-go-camp/ninth/dto"
	"time"

	"github.com/silenceper/pool"
	"net"
	"sync"
)

type toyProtocolClient struct {
	connPool pool.Pool
	m        sync.Map
	invoker  Proxy
}

type Proxy interface {
	Invoke(ctx context.Context, req *dto.Request) (*dto.Response, error)
}

func NewClient(address string) (*toyProtocolClient, error) {
	// Create a connection pool: Initialize the number of connections to 5, the maximum idle connection is 20, and the maximum concurrent connection is 30
	poolConfig := &pool.Config{
		InitialCap: 5,
		MaxIdle:    20,
		MaxCap:     30,
		Factory:    func() (interface{}, error) { return net.Dial("tcp", address) },
		Close:      func(v interface{}) error { return v.(net.Conn).Close() },
		// Ping:       ping,
		// The maximum idle time of the connection, the connection exceeding this time will be closed, which can avoid the problem of automatic failure when connecting to EOF when idle
		IdleTimeout: 15 * time.Second,
	}
	connPool, err := pool.NewChannelPool(poolConfig)
	if err != nil {
		return nil, err
	}
	res := &toyProtocolClient{
		connPool: connPool,
	}
	return res, nil
}

func (c *toyProtocolClient) Invoke(ctx context.Context, req *dto.Request) (*dto.Response, error) {
	cn, err := c.connPool.Get()
	if err != nil {
		return nil, errors.New("could not open new connection")
	}
	// put back
	defer c.connPool.Put(cn)
	bs, _ := json.Marshal(req)

	encode := EncodeMsg(bs)
	_, err = cn.(net.Conn).Write(encode)
	if err != nil {
		return nil, err
	}

	ch := make(chan *dto.Response)
	c.m.Store(req.MessageId, ch)
	bs, err = ReadMsg(cn.(net.Conn))
	if err != nil {
		fmt.Printf("could not read response: %v", err)
		return nil, err
	}

	resp := &dto.Response{}
	err = json.Unmarshal(bs, resp)
	if resp.MessageId <= 0 {
		fmt.Printf("invalid message id: %d", resp.MessageId)
	}
	return resp, nil
}

func (c *toyProtocolClient) invoke(ctx context.Context, req *dto.Request) (chan *dto.Response, error) {
	cn, err := c.connPool.Get()
	if err != nil {
		return nil, errors.New("could not open new connection")
	}
	// put back
	defer c.connPool.Put(cn)
	bs, _ := json.Marshal(req)

	encode := EncodeMsg(bs)
	_, err = cn.(net.Conn).Write(encode)
	if err != nil {
		return nil, err
	}

	ch := make(chan *dto.Response)
	c.m.Store(req.MessageId, ch)
	go func() {
		bs, err = ReadMsg(cn.(net.Conn))
		if err != nil {
			fmt.Printf("could not read response: %v", err)
			return
		}

		resp := &dto.Response{}
		err = json.Unmarshal(bs, resp)
		if resp.MessageId <= 0 {
			fmt.Printf("invalid message id: %d", resp.MessageId)
			return
		}
		resCh, ok := c.m.LoadAndDelete(resp.MessageId)
		if !ok {
			fmt.Printf("could not find the invocation")
			return
		}
		resCh.(chan *dto.Response) <- resp
	}()
	return ch, nil
}

type InjectContextFilterBuilder struct {
	Reader func(ctx context.Context) map[string]string
}

func (i *InjectContextFilterBuilder) Build(next Filter) Filter {
	return func(ctx context.Context, req *dto.Request) (*dto.Response, error) {
		fmt.Printf("this is meta injection filter\n")
		req.Meta = i.Reader(ctx)
		return next(ctx, req)
	}
}
