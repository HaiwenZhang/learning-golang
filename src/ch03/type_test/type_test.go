package type_test

import "testing"

func TestImplicit(t *testing.T) {
	var a int = 1
	var b int64

	b = int64(a) // 不支持隐式类型转换
	t.Log(a, b)
}

func TestPoint(t *testing.T) {
	a := 1
	aPtr := &a
	t.Log(a, aPtr)
	t.Logf("%T %T", a, aPtr)
}

func TestString(t *testing.T) {
	var s string
	t.Log("*" + s + "*")
}
