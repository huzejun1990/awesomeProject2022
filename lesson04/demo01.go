package main

import "fmt"

/*
  函数是基本的代码块，用于执行一个任务
  GO 语言最少有个 main() 函数
  你可以通过函数来划分不同功能，逻辑上每个函数执行的是指定的任务
  函数声明告诉了编译器函数的名称，返回类型、和参数
*/
func main() {
	fmt.Println("hello,world")

	// 调用函数 函数名()
	fmt.Println(add(1, 2))
}

/*func 函数名(参数,参数 ...) 函数调用后的返回值  {
	函数体： 执行一段代码
	return 返回结果
}
*/

func add(a, b int) int {
	c := a + b
	return c
}
