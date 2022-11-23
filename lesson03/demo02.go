package main

import "fmt"

// if 如果。。。否则 else
func main() {

	// 分数
	var score int = 92

	// a b c d
	if score >= 90 && score <= 100 {
		fmt.Println("A")
	} else if score >= 80 && score < 90 {
		fmt.Println("B")
	} else if score >= 70 && score < 80 {
		fmt.Println("C")
	} else if score >= 60 && score < 70 {
		fmt.Println("D")
	} else {
		fmt.Println("不及格")
	}

}
