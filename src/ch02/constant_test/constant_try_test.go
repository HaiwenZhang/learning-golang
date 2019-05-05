package constant_test

import "testing"

const (
	Monday = iota + 1
	Tuesday
	Wednesday
)

const (
	Readable = 1 << iota
	Writable
	Executable
)

func TestConstantTryTest(t *testing.T) {
	t.Log(Monday, Tuesday)
}

func TestConstantTryTest1(t *testing.T) {
	a := 7
	t.Log(a&Readable, a&Writable, a&Executable)
}
