package concurrency

import (
	"fmt"
	"sync"
)

func Once()  {

}

// NeedOnce 用于测试 sync.Once 在方法之间传递
func NeedOnce(once sync.Once) {
	once.Do(func() {
		fmt.Println("calling NeedOnce")
	})
}


// 使用Once，要么调用者自己决定
// 要么方法内部偷偷用了

// NeedOncePrt 用于测试 sync.Once 的指针在方法之间传递
func NeedOncePrt(once *sync.Once) {
	once.Do(func() {
		fmt.Println("calling NeedOncePtr")
	})
}


