package main

import "fmt"

// 类型转换
// 转换后的变量 := 要转换的类型（变量）
// 备注：整型是不能转换为bool类型的
func main() {

	// bool
	a := 3   // int
	b := 5.0 //float64

	// 需求：将int类型的a 转换为 float64类型 类型转换
	c := float64(a)
	d := int(b)

	// bool 整型是不能转换为 bool类型的
	//e := bool(a)

	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b)
	fmt.Printf("%T\n", c)
	fmt.Printf("%T\n", d)
	//fmt.Printf("%T\n",e)

}
