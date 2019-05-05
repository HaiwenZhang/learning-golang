package main // 包，表明代码所在的模块

import "fmt" // 引入代码依赖

// 功能实现
func main() {
	fmt.Println("Hello world")
}

/*
	应用程序的入口：
	    1. 必须是main包： package main
		2. 必须是main方法： func  main()
		3. 文件名不一定是main.go
	与其他语言的差异：
		1. main函数不支持传入参数
		2. 在程序中直接通过os.Args获取命令行参数


*/
