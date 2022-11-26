package main

import "fmt"

func main() {

	// Slice[T]
	type Slice[T int | float64 | float32] []T

	var a Slice[int] = []int{1, 2, 3}
	fmt.Println(a)
	fmt.Printf("Type Name: %T", a) // Type Name: main.Slice[int][1 2 3]
	var b Slice[float32] = []float32{1, 2, 3}
	fmt.Println(b)
	fmt.Printf("Type Name: %T", b)
	var c Slice[float64] = []float64{1, 2, 3}
	fmt.Println(c)
	fmt.Printf("Type Name: %T", c)

	// x 错误。因为变量a的类型为Slice[int],b的类型为Slice[float32],两者类型不同
	//a = b
	//var d Silce[string] = []string{"1","2","3"}
	//fmt.Println(d)
	//fmt.Printf("Type Name: %T",b)
	//var d Silce[T] = []int{1,2,3}

	type MyMap[KEY int | string, VALUE float32 | float64] map[KEY]VALUE

	var m1 MyMap[string, float64] = map[string]float64{
		"go":   9.9,
		"java": 9.0,
	}
	fmt.Println(m1)

	//type Wow[T int | string] int
	//
	//var a Wow[int] = 123	//编译正确
	//var b Wow[string] = 123	// 编译正确
	//var c Wow[string] = "hello"	// 编译错误，因为 "hello"不能赋值给 底层类型int

}
