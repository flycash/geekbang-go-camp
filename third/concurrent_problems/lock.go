package concurrent_problems

import "sync"

var l = sync.RWMutex{}

func RecursiveA()  {
	l.Lock()
	defer l.Unlock()
	RecursiveB()
}

func RecursiveB()  {
	RecursiveC()
}

func RecursiveC()  {
	l.Lock()
	defer l.Unlock()
	RecursiveA()
}

var data = map[string]string{}
func ReadLockToWriteLock(key string) string {
	l.RLock()
	defer l.RUnlock()
	value, ok := data[key]
	if !ok {
		l.Lock()
		data[key] = key + "_123"
		l.Unlock()
	}
	return value
}
