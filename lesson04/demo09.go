package main

import "fmt"

// defer 关闭操作
func main() {
	/*	f("1")
		fmt.Println("2")
		defer f("3")	// 会被延迟到最后执行
		fmt.Println("4")
		defer f("5")	// 会被延迟到最后执行
		fmt.Println("6")
		defer f("7")	// 会被延迟到最后执行
		fmt.Println("8")*/

	a := 10
	fmt.Println("a=", a)
	defer f(a) // 参数就已经传递进去了，在最后执行
	a++
	fmt.Println("end a=", a)
}

/*func f(s string)  {
	fmt.Println("函数里面的a=",s)
}
*/
func f(s int) {
	fmt.Println("函数里面的a=", s)
}
