package easy

var packageVariable string

func SomeLogic(input string) {
	packageVariable = input[0:4]
}

func OtherLogic(input string) {
	packageVariable = input[0:2]
}
