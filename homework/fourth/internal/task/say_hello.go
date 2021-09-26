package task

import (
	"fmt"
	"github.com/gotomicro/ego/task/ejob"
)
// Hello go run main.go --config=configs/local.toml --job=say_hello
func Hello(ctx ejob.Context) error {
	fmt.Printf("hello, world")
	return nil
}
