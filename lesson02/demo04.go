package main

import "fmt"

func main() {

	// := 自动推导
	var name string
	//name := "dream" // no new variables on left side of :=
	age := 18

	fmt.Println(name, age)
	fmt.Printf("%T,%T", name, age)
}
