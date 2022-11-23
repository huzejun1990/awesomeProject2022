package main

import "fmt"

func main() {

	var a int = 11
	var b int = 10

	// == 等于	= 赋值
	// 关系运算符 结果都是 bool
	fmt.Println(a == b) // 等于
	fmt.Println(a != b) // 不等于
	fmt.Println(a > b)
	fmt.Println(a < b)
	fmt.Println(a >= b)
	fmt.Println(a <= b)

	// 判断 if 如果 。。。结果
	if a > b {
		// 执行一些 a > b 的操作
	} else { // else 否则

	}
}
