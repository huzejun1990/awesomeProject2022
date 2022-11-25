package main

import "fmt"

func main() {
	r1 := add(1, 2)
	fmt.Println(r1)

	r2 := oper(3, 4, add)
	fmt.Println(r2)

	r3 := oper(8, 4, sub)
	fmt.Println(r3)

	r4 := oper(8, 4, func(a int, b int) int {
		if b == 0 {
			fmt.Println("除数不能为0")
			return 0
		}
		return a / b
	})
	fmt.Println(r4)

}

// 高阶函数
func oper(a, b int, fun func(int, int) int) int {
	r := fun(a, b)
	return r
}

func add(a, b int) int {
	return a + b
}

func sub(a, b int) int {
	return a - b
}
