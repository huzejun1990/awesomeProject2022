package main

import "fmt"

func main() {

	var x int
	var y float64

	//定义了两个变量，使用键盘来录入这两个变量
	//fmt.Println()		//打印并换行
	//fmt.Printf()		//格式化输出
	//fmt.Print()		// 打印输出

	fmt.Println("请输入两个数：1、整数，2、浮点数：")
	// 变量取地址 &变量
	// 指针、地址来修改和操作变量
	// Scanln 阻塞等待你的键盘输入
	fmt.Scanln(&x, &y)
	fmt.Println("x", x)
	fmt.Println("y", y)
	//fmt.Scanln()	// 接收输入 Scan |
	//fmt.Scan()		// 接收输入 格式化输入 作业
	//fmt.Scanf()		// 接收输入 作业

}
