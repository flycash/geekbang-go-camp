package fourth

import (
	"fmt"
	"github.com/stretchr/testify/suite"
	"strings"
	"testing"
)

//
func BenchmarkFib(b *testing.B) {
	Fib(40)
}

// go test `-bench=BenchmarkStringJoin1 `-v `-benchmem
func BenchmarkStringJoin1(b *testing.B) {
	// b.ReportAllocs() 等效于 -benchmem
	input := []string{"Hello", "World"}
	for i := 0; i < b.N; i++ {
		result := strings.Join(input, " ")
		if result != "Hello World" {
			b.Error("Unexpected result: " + result)
		}
	}
}

func BenchmarkStringJoin2(b *testing.B) {
	// b.ReportAllocs() 等效于 -benchmem
	// b.ReportAllocs()
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

type FibTestSuite struct {
	suite.Suite
}

func (s *FibTestSuite) TearDownTest() {
	fmt.Println("TearDownTest")
}

func (s *FibTestSuite) TearDownSuite() {
	fmt.Println("TearDownSuite")
}

func (s *FibTestSuite) SetupTest() {
	fmt.Println("SetupTest")
}

func (s *FibTestSuite) SetupSuite() {
	fmt.Println("SetupSuite")
}

func (s *FibTestSuite) BeforeTest(suiteName, testName string) {
	fmt.Println("BeforeTest")
}

func (s *FibTestSuite) AfterTest(suiteName, testName string) {
	fmt.Println("AfterTest")
}

func (s *FibTestSuite) TestFib() {
	res := Fib(3)
	println(res)
}

func (s *FibTestSuite) TestFibCopy() {
	res := Fib(3)
	if res != 2 {
		s.T().Fatal("incorrect")
	}
}

func TestFib(t *testing.T) {
	suite.Run(t, new(FibTestSuite))
}

func ExampleFib() {
	// pass a number
	Fib(3)
	// Output: 3
}
