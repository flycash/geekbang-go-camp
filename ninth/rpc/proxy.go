package rpc

import (
	"context"
	"encoding/json"
	"geekbang/geekbang-go-camp/ninth/dto"
	"reflect"
	"sync/atomic"
)

var messageId uint64 = 1

func InitToyProtocolProxy(address string, val Service, ftls ...FilterChain) error {
	c, err := NewClient(address)
	p := NewFilterProxy(c, ftls...)
	if err != nil {
		return err
	}
	setFuncField(val, p)
	return nil
}

func setFuncField(val Service, c Proxy) {

	v := reflect.ValueOf(val)
	ele := v.Elem()
	t := ele.Type()

	numField := t.NumField()
	for i := 0; i < numField; i++ {
		field := t.Field(i)
		fieldValue := ele.Field(i)
		if fieldValue.CanSet() {
			fn := func(args []reflect.Value) (results []reflect.Value) {
				in := args[1].Interface()
				out := reflect.New(field.Type.Out(0).Elem()).Interface()
				inData, err := json.Marshal(in)

				req := &dto.Request{
					ServiceName: val.ServiceName(),
					Method:      field.Name,
					Data:        inData,
					MessageId:   atomic.AddUint64(&messageId, 1),
				}

				if err != nil {
					return []reflect.Value{reflect.ValueOf(out), reflect.ValueOf(err)}
				}

				resp, err := c.Invoke(args[0].Interface().(context.Context), req)
				if err != nil {
					return []reflect.Value{reflect.ValueOf(out), reflect.ValueOf(err)}
				}

				err = json.Unmarshal(resp.Data, out)

				if err != nil {
					return []reflect.Value{reflect.ValueOf(out), reflect.ValueOf(err)}
				}

				return []reflect.Value{reflect.ValueOf(out), reflect.Zero(reflect.TypeOf(new(error)).Elem())}
			}
			fieldValue.Set(reflect.MakeFunc(field.Type, fn))
		}
	}
}
