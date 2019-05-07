package main

import (
	"ch09/series"
	"fmt"
)

func main() {
	fibList, _ := series.GetFibonacci(10)
	fmt.Println(fibList)
}
