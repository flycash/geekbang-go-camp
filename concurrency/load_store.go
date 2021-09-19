package concurrency

import "fmt"

type connection struct {

}

func (c *connection) send() error {
	fmt.Println("send...")
	return nil
}

func (c *connection) Close() error {
	fmt.Println("closing...")
	return nil
}

