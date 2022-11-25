package main

import "fmt"

//switch语句
func main() {

	var score int = 90

	// 匹配 case
	switch score {
	case 90:
		fmt.Println("A")
		fallthrough // 很少用
	case 80:
		if score == 90 {
			break
		}
		fmt.Println("B")
		fallthrough
	case 50, 60, 70:
		fmt.Println("C")
		fallthrough
	default:
		fmt.Println("D")

	}
	// switch 默认的条件 bool = true
	switch {
	case false:
		fmt.Println("false")
	case true:
		fmt.Println("true")
	default:
		fmt.Println("其他")
	}

}
