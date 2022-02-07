package rpc

import (
	"context"
	"encoding/json"
	"geekbang/geekbang-go-camp/ninth/dto"
	"reflect"
)

type Service interface {
	ServiceName() string
}

type ProviderProxy interface {
	Invoke(req *dto.Request) *dto.Response
}
type reflectionStub struct {
	s     Service
	value reflect.Value
}

func (s *reflectionStub) invoke(ctx context.Context, methodName string, data []byte) ([]byte, error) {
	method := s.value.MethodByName(methodName)
	inType := method.Type().In(1)
	in := reflect.New(inType.Elem())
	err := json.Unmarshal(data, in.Interface())
	if err != nil {
		return nil, err
	}
	res := method.Call([]reflect.Value{reflect.ValueOf(ctx), in})
	if len(res) > 1 && !res[1].IsZero() {
		return nil, res[1].Interface().(error)
	}
	return json.Marshal(res[0].Interface())
}
