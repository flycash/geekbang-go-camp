package sub

import (
	"geekbang/geekbang-go-camp/concurrency"
	"testing"
)


func TestSingleton(t *testing.T)  {
	//s := concurrency.GetSingleInstance()
	//s.Single()

	var s = concurrency.GetSingleInstance()
	s.Single()
}
