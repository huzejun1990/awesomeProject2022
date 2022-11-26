// @Author huzejun 2022/11/27 0:00:00
package main

import "fmt"

type MySlice[T int | float64] []T

// 给泛型添加方法
func (s MySlice[T]) Sum() T {
	var sum T
	for _, v := range s {
		sum += v
	}
	return sum
}

// 泛型函数
/*func Add(a int,b int) int {
	return a+b
}*/

func Add[T int | float64 | string](a T, b T) T {
	return a + b
}

func main() {

	fmt.Println(Add[int](1, 2))
	fmt.Println(Add(1, 2))
	fmt.Println(Add[string]("1", "hehe"))
	fmt.Println(Add("1", "hehe"))
	fmt.Println(Add[float64](1.2, 3.1415))
	fmt.Println(Add(1.2, 3.1415))
	/*	var s MySlice[int] = []int{1,2,3,4}
		fmt.Println(s.Sum())

		var f MySlice[float64] = []float64{1.0,2.2,3.3,4.5}
		fmt.Println(f.Sum())*/
}
