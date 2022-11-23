package main

import "fmt"

func main() {
	var a int = 10
	var b int = 3
	//var c int

	// + - * / % ++ --
	fmt.Println(a + b) // 13
	fmt.Println(a - b) // 7
	fmt.Println(a * b) // 30
	fmt.Println(a % b) // 1
	a++                // a = a+1	// a = a+1
	fmt.Println(a)     // 11

	a = 100
	a--            // a = a - 1
	fmt.Println(a) // 10
}
