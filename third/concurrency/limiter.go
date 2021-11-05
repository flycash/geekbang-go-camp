package concurrency

import (
	"sync/atomic"
)

type Limiter struct {
	// 当前处理请求的上限
	limit int32
	// 处理请求逻辑
	handler func(req interface{}) interface{}

	cnt int32
}

// Reject bool 返回值表示究竟有没有执行
func (l *Limiter) Reject(req interface{}) (interface{}, bool) {

	// 进来，二话不说，我就直接先给你加了，就是说，我分配了一个位置给你
	cnt := atomic.AddInt32(&l.cnt, 1)
	defer atomic.AddInt32(&l.cnt, -1)
	// 我发现我超出了上限
	if cnt > l.limit {
		return nil, false
	}
	res := l.handler(req)
	return res, true
}

// v3
//type Limiter struct {
//	// 当前处理请求的上限
//	limit int
//	// 处理请求逻辑
//	handler func(req interface{}) interface{}
//
//	mutex sync.RWMutex
//	cnt int
//}
//
//// Reject bool 返回值表示究竟有没有执行
//func (l *Limiter) Reject(req interface{}) (interface{}, bool) {
//	l.mutex.RLock()
//	if l.cnt >= l.limit {
//		l.mutex.RUnlock()
//		return nil, false
//	}
//	l.mutex.Lock()
//	if l.cnt >= l.limit {
//		l.mutex.Unlock()
//		return nil, false
//	}
//
//	l.cnt ++
//	l.mutex.Unlock()
//	res := l.handler(req)
//	l.mutex.Lock()
//	defer l.mutex.Unlock()
//	l.cnt --
//	return res, true
//}

// v2
//type Limiter struct {
//	// 当前处理请求的上限
//	limit int
//	// 处理请求逻辑
//	handler func(req interface{}) interface{}
//
//	mutex sync.Mutex
//	cnt int
//}
//
//// Reject bool 返回值表示究竟有没有执行
//func (l *Limiter) Reject(req interface{}) (interface{}, bool) {
//	l.mutex.Lock()
//	if l.cnt < l.limit {
//		l.cnt ++
//		l.mutex.Unlock()
//		res := l.handler(req)
//		l.mutex.Lock()
//		defer l.mutex.Unlock()
//		l.cnt --
//		return res, true
//	}
//	l.mutex.Unlock()
//	return nil, false
//}


// v1
//type Limiter struct {
//	// 当前处理请求的上限
//	limit int
//	// 处理请求逻辑
//	handler func(req interface{}) interface{}
//
//	mutex sync.Mutex
//}
//
//// Reject bool 返回值表示究竟有没有执行
//func (l *Limiter) Reject(req interface{}) (interface{}, bool) {
//	l.mutex.Lock()
//	defer l.mutex.Unlock()
//	res := l.handler(req)
//	return res, true
//}