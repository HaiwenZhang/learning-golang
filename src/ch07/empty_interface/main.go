package main

import "fmt"

func main() {
	DoSomething(10)
	DoSomething("10")
}

func DoSomething(p interface{}) {
	// if i, ok := p.(int); ok {
	// 	fmt.Println("Integer ", i)
	// 	return
	// }
	// if s, ok := p.(string); ok {
	// 	fmt.Println("string ", s)
	// 	return
	// }
	// fmt.Println("Unknow Type")
	switch v := p.(type) {
	case int:
		fmt.Println("Integer ", v)
	case string:
		fmt.Println("string ", v)
	default:
		fmt.Println("Unknow Type")
	}
}
