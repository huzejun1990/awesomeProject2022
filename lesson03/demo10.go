package main

import "fmt"

func main() {

	str := "hello,dream"
	fmt.Println(str)

	// 获取字符串的长度 len
	fmt.Println("字符串的长度为：", len(str))

	// 获取指定的字节
	fmt.Println("字节打印：", str[0])
	fmt.Printf("%c", str[2])

	// for
	for i := 0; i < len(str); i++ {
		fmt.Printf("%c\n", str[i])
	}

	// for range循环，遍历数组、切片...
	//返回下标和对应的值，使用这个值就可以了
	for i, v := range str {
		fmt.Print(i)
		fmt.Printf("%c", v)
	}

	// string 不能修改
	//str[2] = 'A'
	//fmt.Println(str[2])
}
