package string_test

import "testing"

func TestString(t *testing.T) {
	var s string
	s = "hello"
	t.Log(len(s))
	s = "\xE4\xB8\xA5"
	t.Log(s)
}
