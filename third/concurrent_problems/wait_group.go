package concurrent_problems

import (
	"fmt"
	"sync"
)

func WrongAddFunction()  {
	wg := sync.WaitGroup{}
	result := 0
	for i := 0; i < 10; i++ {
		go func(delta int) {
			wg.Add(1)
			result += delta
			defer wg.Done()
		}(i)
	}

	wg.Wait()
	fmt.Println(result)
}

func WrongLock() {
	lock := sync.Mutex{}
	go func() {
		lock.Lock()
	}()
	lock.Unlock()
}
