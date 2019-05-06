package main

import "fmt"

func main() {
	s := "中"
	c := []rune(s)
	fmt.Println(len(s))
	fmt.Println(len(c))
	fmt.Printf("中 unicode %x", c[0])
	fmt.Printf("中 UTF8 %x", s)
}
