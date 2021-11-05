package sub

import (
	"geekbang/geekbang-go-camp/third/concurrency"
	"testing"
)


func TestSingleton(t *testing.T)  {
	//s := concurrency.GetSingleInstance()
	//s.Single()

	var s = concurrency.GetSingleInstance()
	s.Single()
}
