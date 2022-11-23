package main

import "fmt"

func main() {
	const URL string = "www.baidu.com" // 显示定义
	const URL2 = "www.baidu.com"       // 隐式定义

	const a, b, c = 3.14, "deam", false

	fmt.Println(URL)
	fmt.Println(URL2)
	fmt.Println(a, b, c)
}
