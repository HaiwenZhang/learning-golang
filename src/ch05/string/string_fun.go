package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "A,B,C"
	parts := strings.Split(s, ",")
	fmt.Println(parts)
}
