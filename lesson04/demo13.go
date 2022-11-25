package main

import "fmt"

/*
一个外层函数中，有内存函数，该内存函数中，会操作外层函数的局部变量
并且该外层函数的返回值就是这个内层函数，
这个内存函数和外层函数的局部变量，统称为闭包结构

局部变量的生命周期就会发生改变，正常的局部变量会随着函数的调用而创建，随函数的结束而销毁
但是闭包结构中的外层函数的局部变量平区不会随着外层函数的结束而销毁，因为内层函数还在继续使用
*/

func main() {

	r1 := increment()
	fmt.Println(r1)

	v1 := r1()
	fmt.Println(v1)
	v2 := r1()
	fmt.Println(v2)
	fmt.Println(r1())
	fmt.Println(r1())
	fmt.Println(r1())

	// inclement(), r1 已经销毁了
	r2 := increment()
	v3 := r2()
	fmt.Println(v3)
	fmt.Println(r1())
	fmt.Println(r1())
}

//自增
func increment() func() int {
	//局部变量 i
	i := 0
	// 定义一个匿名函数，给变量自增并返回
	fun := func() int { // 内层函数，没有执行的
		i++
		return i
	}
	return fun
}
