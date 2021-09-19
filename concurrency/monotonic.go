package concurrency

type filter struct {
	// 处理请求逻辑
	handler func(req interface{}) interface{}
	reject int
}

func (f *filter) Handle(req interface{}) (interface{}, bool) {
	if f.reject > 0 {
		return nil, false
	}
	return f.handler(req), true
}

func (f *filter) RejectNewRequest() {
	f.reject = 1
}