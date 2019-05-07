package main

import "fmt"

type Pet struct {
}

func (p *Pet) Speak() {
	fmt.Print("...")
}

func (p *Pet) SpeakTo(host string) {
	p.Speak()
	fmt.Println(" ", host)
}

// type Dog struct {
// 	p *Pet
// }

// func (d *Dog) Speak() {
// 	d.p.Speak()
// }

// func (d *Dog) SpeakTo(host string) {
// 	d.p.SpeakTo(host)
// }

type Dog struct {
	Pet
}

type Code string
type Programmer interface {
	WriteHelloWorld() Code
}
type GoProgrammer struct {
}

func (p *GoProgrammer) WriteHelloWorld() Code {
	return "goprogrammer"
}

type JavaProgrammer struct {
}

func (p *JavaProgrammer) WriteHelloWorld() Code {
	return "Java programmer"
}

func writeFirstProgram(p Programmer) {
	fmt.Printf("%T %v\n", p, p.WriteHelloWorld())
}

func main() {
	dog := new(Dog)
	dog.SpeakTo("Test")

	goProg := new(GoProgrammer)
	javaProg := new(JavaProgrammer)
	writeFirstProgram(goProg)
	writeFirstProgram(javaProg)

}
