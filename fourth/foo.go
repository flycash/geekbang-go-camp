package fourth

type Foo interface {
	Bar(x int) int
}

func SUT(f Foo) {
	res := f.Bar(99)
	println(res)
}
