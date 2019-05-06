package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	defer Clear()

	a, b := returnMultiValues()

	fmt.Printf(string(a))
	fmt.Printf(string(b))
	tsSF := timeSpent(slowFun)
	tsSF(10)

	res := Sum(1, 2, 3, 4, 5, 6, 7)
	fmt.Println(res)
}

func timeSpent(inner func(op int) int) func(op int) int {
	return func(n int) int {
		start := time.Now()
		ret := inner(n)
		fmt.Printf("time spent: ", time.Since(start).Seconds())
		return ret
	}
}

func slowFun(op int) int {
	time.Sleep(time.Second * 1)
	return op
}

func returnMultiValues() (int, int) {
	return rand.Intn(10), rand.Intn(20)
}

func Sum(ops ...int) int {
	ret := 0
	for _, op := range ops {
		ret += op
	}
	return ret
}

func Clear() {
	fmt.Println("clear resources.")
}
