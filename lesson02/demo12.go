package main

import "fmt"

func main() {

	var str string
	str = "Hello,dream"
	fmt.Printf("%T,%s\n", str, str)

	// 单引号 字符，整型-ASCII字符表
	v1 := '中'
	v2 := "A"
	// 编码表 ASCII字符码
	// 扩展：
	// 所有的中国字的编码表：GBK
	// 全世界的编码表：Unicode编码表

	fmt.Printf("%T,%d\n", v1, v1) // int32,65
	fmt.Printf("%T,%s\n", v2, v2) // string,A

	// 字符串连接 +
	fmt.Println("hello" + ",dream")

	//转义字符 \
	fmt.Println("hello\"dream")
	fmt.Println("hello\ndream") // \n 换行
	fmt.Println("hello\tdream") // t 制表符

}
