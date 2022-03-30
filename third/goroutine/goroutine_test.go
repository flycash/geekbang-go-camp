package goroutine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestNewTaskPool(t *testing.T) {
	tp := NewBlockTaskPool(2)
	var wg sync.WaitGroup
	// 这个 wg 没啥意义，就是确保我们的测试可以在所有的 task 都执行完才退出，不用关注
	wg.Add(10)
	for i := 0; i < 10; i++ {
		cnt := i
		tp.Do(func() {
			fmt.Printf("task %d start \n", cnt)
			time.Sleep(time.Second)
			fmt.Printf("task %d end \n", cnt)
			wg.Done()
		})
	}
	// 这里打印 NumGoroutine 会多几个，因为 Go 本身和测试框架本身都会开启必须的 goroutine
	// println(runtime.NumGoroutine())
	wg.Wait()
}

func TestNewCacheTaskPool(t *testing.T) {
	tp := NewCacheBlockTaskPool(2, 10)
	tp.Start()

	var wg sync.WaitGroup
	// 这个 wg 没啥意义，就是确保我们的测试可以在所有的 task 都执行完才退出，不用关注
	wg.Add(20)
	for i := 0; i < 20; i++ {
		cnt := i
		tp.Do(func() {
			fmt.Printf("task %d start \n", cnt)
			time.Sleep(time.Second)
			fmt.Printf("task %d end \n", cnt)
			wg.Done()
		})
	}
	// 这里打印 NumGoroutine 会多几个，因为 Go 本身和测试框架本身都会开启必须的 goroutine
	// println(runtime.NumGoroutine())
	wg.Wait()
}
