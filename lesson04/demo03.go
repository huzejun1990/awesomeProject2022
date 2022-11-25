package main

import "fmt"

func main() {

	// 形参与实参要一一对应，顺序，个数、类型
	max(1, 2)
	fmt.Println(max(1, 2))
}

// max 两个数字比大小
//形式参数: 定义函数时，用来接收外部传入数据的参数，就是形式参数
//实际参数：调用函数时，传给形参的实际数据叫做实际参数
func max(num1, num2 int) int {

	var result int
	if num1 > num2 {
		result = num1
	} else {
		result = num2
	}

	//一个函数定义上有返回值，那么函数中必须使用return语句
	// 返回值
	// 调用处需要使用变量接收该结果
	return result
}
