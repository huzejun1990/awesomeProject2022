package main

import "fmt"

func main() {
	//函数的调用
	printinfo()
	myprint("hehe")

	// 有返回值的函数，我们就需要接收函数的返回值
	c := add2(1, 2)
	myprintnum(c)

	x, y := swap("world", "hello")
	fmt.Println(x, y)
}

//无参无返回值函数
func printinfo() {
	fmt.Println("printinfo")
}

//有一个参数的值函数、、、
func myprint(msg string) {
	fmt.Println(msg)
}

func myprintnum(num int) {
	fmt.Println(num)
}

//有一个返回值的函数
//两个参数的值函数
func add2(a, b int) int {
	c := a + b
	return c
}

//无参无返回值函数
//有多个返回值函数
func swap(x, y string) (string, string) {
	return y, x
}
