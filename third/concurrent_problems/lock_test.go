package concurrent_problems

import "testing"

func TestRecursiveA(t *testing.T) {
	RecursiveA()
}

func TestReadLockToWriteLock(t *testing.T) {
	ReadLockToWriteLock("mykey")
}