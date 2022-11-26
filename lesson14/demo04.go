// @Author huzejun 2022/11/27 0:15:00
package main

import "fmt"

type intAAA int8

// 如果类型太多了如何处理？这时候我们就可以用不用 自定义泛型类型
type MyInt interface {
	int | ~int8 | int16 | int32 | int64
}

func GetMaxNum[T MyInt | float64](a, b T) T {
	if a > b {
		return a
	}
	return b
}

func main() {

	var a intAAA = 10
	var b intAAA = 20
	fmt.Println(GetMaxNum(a, b))
	fmt.Println(GetMaxNum[intAAA](a, b))

}
