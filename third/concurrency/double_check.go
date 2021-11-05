package concurrency

import "sync"

type SafeMap struct {
	m map[string]interface{}
	mutex sync.RWMutex
}

// LoadOrStore loaded 代表是返回老的对象，还是返回了新的对象
func (s *SafeMap) LoadOrStore(key string, newVale interface{}) (val interface{}, loaded bool) {
	s.mutex.RLock()
	val, ok := s.m[key]
	s.mutex.RUnlock()
	if ok {
		return val, true
	}
	s.mutex.Lock()
	defer s.mutex.Unlock()
	val, ok = s.m[key]
	if ok {
		return val, true
	}
	s.m[key] = newVale
	return newVale, false
}

type valProvider func() interface{}

func (s *SafeMap) LoadOrStoreHeavy(key string, p valProvider) (val interface{}, loaded bool) {
	s.mutex.RLock()
	val, ok := s.m[key]
	s.mutex.RUnlock()
	if ok {
		return val, true
	}
	s.mutex.Lock()
	defer s.mutex.Unlock()
	val, ok = s.m[key]
	if ok {
		return val, true
	}
	newVale := p()
	s.m[key] = newVale
	return newVale, false
}

func (s *SafeMap)CheckAndDoSomething() {
	s.mutex.Lock()
	// check and do something
	s.mutex.Unlock()
}

func (s *SafeMap) CheckAndDoSomething1() {
	s.mutex.RLock()
	// check 第一次检查
	s.mutex.RUnlock()

	s.mutex.Lock()
	// check and doSomething
	defer s.mutex.Unlock()
}