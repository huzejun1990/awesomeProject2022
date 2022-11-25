package main

import "fmt"

func main() {

	a := false
	switch a {

	case false:
		fmt.Println("1、case条件为false")
		fallthrough // case穿透的，不管下一个条件满不满足，都会执行
	case true:
		if a == false {
			break // 终于case穿透
		}

		fmt.Println("2、case条件为true")

	}
}
