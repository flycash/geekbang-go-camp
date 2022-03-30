package ctx

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestTimeoutExample(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	bsChan := make(chan struct{})
	go func() {
		slowBusiness()
		bsChan <- struct{}{}
	}()
	select {
	case <-ctx.Done():
		fmt.Println("timeout")
	case <-bsChan:
		fmt.Println("business end")
	}
}

func slowBusiness() {
	time.Sleep(2 * time.Second)
}
