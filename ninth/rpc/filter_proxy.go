package rpc

import (
	"context"
	"geekbang/geekbang-go-camp/ninth/dto"
)

type FilterChain func(next Filter) Filter

type Filter func(ctx context.Context, req *dto.Request) (*dto.Response, error)

type filtersProxy struct {
	root Filter
}

func NewFilterProxy(delegate Proxy, fts ...FilterChain) Proxy {
	root := delegate.Invoke
	for i := len(fts) - 1; i >= 0; i-- {
		root = fts[i](root)
	}
	return &filtersProxy{
		root: root,
	}
}

func (f *filtersProxy) Invoke(ctx context.Context, req *dto.Request) (*dto.Response, error) {
	return f.root(ctx, req)
}
