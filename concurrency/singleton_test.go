package concurrency

import "testing"

func TestGetSingleInstance(t *testing.T) {
	ins := GetSingleInstance()
	ins.Single()
}
