package concurrent_problems

import "sync"

type Receiver struct {
	lock sync.Mutex
}

func (r Receiver) Lock() {
	r.lock.Lock()
}

func (r Receiver) Unlock() {
	r.lock.Unlock()
}
