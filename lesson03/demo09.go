package main

import "fmt"

func main() {

	//break 结束当前整个循环
	for i := 0; i < 10; i++ {
		if i == 5 {
			break
		}
		fmt.Println(i)
	}

	// continue 结束单次循环，整个循环还是会继续执行下去
	fmt.Println("===============")
	for i := 0; i < 10; i++ {
		if i == 5 {
			continue
		}
		fmt.Println(i)
	}
}
