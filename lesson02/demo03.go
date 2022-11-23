package main

import "fmt"

func main() {
	// var 定义变量，如果没有赋值，变量应是这个类型的默认值
	var (
		name string
		age  int
		addr string
	)
	name = "dream"
	age = 30
	addr = "china"

	// string默认值： 空
	// int 默认值 ：0
	fmt.Println(name, age, addr)
}
