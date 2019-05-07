package main

import (
	"fmt"
	"time"
)

func main() {
	Select()
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

func Select() {
	select {
	case ret := <-AnsycService():
		fmt.Println(ret)
	case <-time.After(time.Millisecond * 50):
		fmt.Println("time out")
	}
}
