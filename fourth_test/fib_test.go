package fourth

import (
	"strings"
	"testing"
)
//
func BenchmarkFib(b *testing.B) {
	Fib(40)
}

// go test `-bench=BenchmarkStringJoin1 `-v `-benchmem
func BenchmarkStringJoin1(b *testing.B) {
	//b.ReportAllocs() 等效于 -benchmem
	input := []string{"Hello", "World"}
	for i := 0; i < b.N; i++ {
		result := strings.Join(input, " ")
		if result != "Hello World" {
			b.Error("Unexpected result: " + result)
		}
	}
}

func BenchmarkStringJoin2(b *testing.B) {
	//b.ReportAllocs() 等效于 -benchmem
	//b.ReportAllocs()
	input := []string{"Hello", "World"}
	join := func(strs []string, delim string) string {
		if len(strs) == 2 {
			return strs[0] + delim + strs[1]
		}
		return ""
	}
	for i := 0; i < b.N; i++ {
		result := join(input, " ")
		if result != "Hello World" {
			b.Error("Unexpected result: " + result)
		}
	}
}
