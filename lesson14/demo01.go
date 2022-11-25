package main

import "fmt"

// Go泛型编程讲解
func main() {

	strs := []string{"dream", "hello,world"}
	is := []int{1, 2}
	fs := []float64{1, 2}
	//printStringArray(is)
	printArray(is)
	printArray(fs)
	printArray(strs)
}

//** 我们不限定他的类型，让调用者自己去定义类型 **
// []T 形式类型 实际类型
/*
** 内置的泛型类型any和comparable**

- **any: ** 表示go里面所有的内置基本类型，等价于‘interface{}’
- ** comparable:** 表示go里面所有内置的可比较类型：`int、uint、float、bool、struct、指针`等一切可以比较的类型

*/
func printArray[T any](arr []T) {
	//类型断言，X.(T) 其实就是判断 T 是否实现了 X 接口，如果实现了，就把 X 接口类型具体化为 T 类型
	for _, a := range arr {
		fmt.Println(a)
	}
}

func printIntArray(arr interface{}) {
	//类型断言，X.(T) 其实就是判断 T 是否实现了 X 接口，如果实现了，就把 X 接口类型具体化为 T 类型
	for _, a := range arr.([]int) {
		fmt.Println(a)
	}
}
