package main

import "fmt"

// 目标：使用for循环打印99乘法表
func main() {
	// j 1 2 3 4 5 6
	for j := 1; j <= 9; j++ {
		for i := 1; i <= j; i++ {
			fmt.Printf("%d*%d=%d \t", i, j, i*9)
		}
		fmt.Println()
	}

}
