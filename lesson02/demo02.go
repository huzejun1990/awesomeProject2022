package main

import "fmt"

// 全局变量
var name string = "dream"

func main() {

	// = 赋值，应是将右边的值，赋给左边的变量
	// name 变量
	//局部变量
	var age int = 18
	var name string = "lisi"

	fmt.Println(name, age)
}

func aaa() {
	fmt.Println(name)
}
