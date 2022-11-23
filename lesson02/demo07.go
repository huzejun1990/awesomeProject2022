package main

import "fmt"

/*
定义一个test函数，它会返回两个int类型的值，每次调用 将会返回 100 和 200 两个数值
我们用这个函数理解 _ 这个是匿名变量。
*/

func test() (int, int) {
	return 100, 200
}

func main() {
	//对象：User
	a, _ := test()
	_, b := test()
	fmt.Println(a, b)

}
