package main

import "fmt"

type Programmer interface {
	WriterHelloWorld() string
}

type GoProgrammer struct {
}

func (go1 *GoProgrammer) WriterHelloWorld() string {
	return "fmt.Println(\"Hello world\")"
}

func main() {
	var p Programmer
	p = new(GoProgrammer)
	fmt.Println(p.WriterHelloWorld())
}
