package main

import "fmt"

func main() {
	//定义一个整型
	// byte uint8
	// rune int32
	// int int64

	var age int = 200
	fmt.Printf("%T,%d\n", age, age)

	//定义一个浮点型
	//默认6位小数打印 3.140000
	var money float64 = 3.14
	fmt.Printf("%T,%f\n", money, money)

	// 数据类型：内存空间分配的大小是不同的

	// float64 尽量使用 float64 来定义浮点类型的小数
	var num2 float64 = -123.0000901
	var num1 float32 = float32(num2)
	fmt.Println("num1 = ", num1, "num2 = ", num2)

}
