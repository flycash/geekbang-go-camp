package goroutine

import (
	"fmt"
	"go.uber.org/atomic"
)

// BlockTaskPool 会阻塞任务的提交。
// 例如我们创建的最大并发数量是 5，
// 那么如果此时已经有五个任务正在执行，第六个提价的时候会被阻塞
type BlockTaskPool struct {
	// 用一个 channel 来控制数量
	concurrent chan struct{}

	// 这个字段不用关注，只是我用来打印开启 goroutine 用的
	cnt atomic.Int32
}

func NewBlockTaskPool(maxCon int) *BlockTaskPool {
	tp := &BlockTaskPool{
		concurrent: make(chan struct{}, maxCon),
	}
	return tp
}

func (tp *BlockTaskPool) Do(action func()) {
	tp.concurrent <- struct{}{}
	go func() {
		tp.cnt.Add(1)
		fmt.Printf("task goroutine num %d \n", tp.cnt.Load())
		action()
		<-tp.concurrent
		tp.cnt.Sub(1)
	}()
}

// CacheBlockTaskPool 带缓存的
// 只要提交的任务数量没有超过 queue 的容量，那么就不会阻塞
// 比如说 NewCacheBlockTaskPool(5, 10) 代表最多有五个任务同时运行，但是还可以再缓存10个任务
type CacheBlockTaskPool struct {
	// 用一个 channel 来控制数量
	concurrent chan struct{}
	queue      chan func()

	// 这个字段不用关注，只是我用来打印开启 goroutine 用的
	cnt atomic.Int32
}

func NewCacheBlockTaskPool(maxCon int, queueSize int) *CacheBlockTaskPool {
	tp := &CacheBlockTaskPool{
		concurrent: make(chan struct{}, maxCon),
		queue:      make(chan func(), queueSize),
	}
	return tp
}

func (tp *CacheBlockTaskPool) Do(action func()) {
	tp.queue <- action
}

func (tp *CacheBlockTaskPool) Start() {
	go func() {
		for {
			tp.concurrent <- struct{}{}
			go func() {
				tp.cnt.Add(1)
				fmt.Printf("task goroutine num %d , queue size: %d \n", tp.cnt.Load(), len(tp.queue))
				tk := <-tp.queue
				tk()
				<-tp.concurrent
				tp.cnt.Sub(1)
			}()
		}
	}()
}

// Close 作为一个思考题，怎么关闭 Start 里面开启的最外围的 goroutine
func (tp *CacheBlockTaskPool) Close() {

}
