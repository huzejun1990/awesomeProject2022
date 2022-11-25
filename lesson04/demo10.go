package main

import "fmt"

// func() 本身就是一个数据类型
func main() {

	// f1 如果不加括号，函数就是一个变量
	// f1()如果为了括号那就成了函数的调用
	fmt.Printf("%T\n", f1) // func() | func(int, int) | func(int, int) int
	/*	fmt.Printf("%T\n",10)	// int
		fmt.Printf("%T\n","hello")	// string
	*/
	// 定义函数类型的变量
	var f5 func(int, int)
	f5 = f1
	fmt.Println(f5)
	fmt.Println(f1)
	f5(1, 2)
}

func f1(a, b int) {
	fmt.Println(a, b)
}
