package fourth

import (
	mock_fourth "geekbang/geekbang-go-camp/fourth/mock_test"
	"github.com/golang/mock/gomock"
	"testing"
)

func TestFoo(t *testing.T) {
	ctrl := gomock.NewController(t)

	// Assert that Bar() is invoked.
	defer ctrl.Finish()

	foo := mock_fourth.NewMockFoo(ctrl)

	// Asserts that the first and only call to Bar() is passed 99.
	// Anything else will fail.
	foo.EXPECT().
		Bar(gomock.Eq(99)).
		Return(101)
	SUT(foo)
}




