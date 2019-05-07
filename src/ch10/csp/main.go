package main

import (
	"fmt"
	"time"
)

func main() {
	multiService()

	fmt.Println("....................")
	TestAnsynService()
}

func service() string {
	time.Sleep(time.Millisecond * 50)
	return "Done"
}

func otherTask() {
	fmt.Println("working on something else")
	time.Sleep(time.Millisecond * 100)
	fmt.Println("Task is done.")
}

func multiService() {
	fmt.Println(service())
	otherTask()
}

func AnsycService() chan string {
	retCh := make(chan string)

	go func() {
		ret := service()
		fmt.Println("returned result")
		retCh <- ret
		fmt.Println("service exited")
	}()

	return retCh
}

func TestAnsynService() {
	retCh := AnsycService()
	otherTask()
	fmt.Println(<-retCh)
}
