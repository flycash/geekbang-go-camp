package concurrent_problems

import (
	"fmt"
	"testing"
	"time"
)

func TestReceiver(t *testing.T) {
	r := Receiver{}
	go func() {
		r.Lock()
		fmt.Println("first lock")
	}()

	go func() {
		r.Lock()
		fmt.Println("second lock")
	}()
	time.Sleep(time.Second)
}