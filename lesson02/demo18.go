package main

import "fmt"

func main() {

	var a int = 21
	var b int

	// = 赋值
	b = a
	fmt.Println(b)

	// += b = a + b = 42
	// 写简化代码的时候，会这样操作
	b += a
	fmt.Println(b)

	// -= /= *= %=
	b = a

	b -= a
	fmt.Println(b)

	b = a

	b *= a
	fmt.Println(b)

	b = a
	b /= a
	fmt.Println(b)

	b = a
	b %= a
	fmt.Println(b)
}
