package main

import "fmt"

func main() {
	f1()
	f2 := f1 //函数本身也是一个变量
	f2()

	// 匿名函数
	f3 := func() {
		fmt.Println("我是f3函数")
	}
	f3()

	// 匿名函数自己调用自己
	func() {
		fmt.Println("我是f4函数")
	}()

	func(a, b int) {
		fmt.Println(a, b)
		fmt.Println("我是f4函数")
	}(1, 2)

	// 匿名函数自己调用自己,可以有参数的
	func(a, b int) {
		fmt.Println(a, b)
		fmt.Println("我是f5函数")
	}(1, 2)

	r1 := func(a, b int) int {
		return a + b
		//fmt.Println("我是f5函数")
	}(1, 2)
	fmt.Println(r1)

}

func f1() {
	fmt.Println("我是f1函数")
}
