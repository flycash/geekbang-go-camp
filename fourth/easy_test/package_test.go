package easy

import (
	"fmt"
	"testing"
)

func TestOtherLogic(t *testing.T) {
	for i := 0; i < 99; i++ {
		ci := i
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			SomeLogic(fmt.Sprintf("OtherLogicTest_%d", ci))
		})
	}
}

func TestSomeLogic(t *testing.T) {
	for i := 0; i < 99; i++ {
		ci := i
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			SomeLogic(fmt.Sprintf("OtherLogicTest_%d", ci))
		})
	}
}
