package concurrency

import (
	"sync"
	"testing"
)

func TestNeedOnce(t *testing.T) {
	var once sync.Once
	// 我们调用两次，希望只输出一次
	// 然而结果是输出了两次
	// 本质上，NeedOnce接收的是一个副本
	NeedOnce(once)
	NeedOnce(once)


	// 在使用指针的时候，只输出了一次
	NeedOncePrt(&once)
	NeedOncePrt(&once)
}
