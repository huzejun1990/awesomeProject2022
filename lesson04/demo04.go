package main

import "fmt"

func main() {
	getSum("hehe", 1, 2, 3, 4, 5, 6, 7, 8, 100)

}

//...可变参数
func getSum(msg string, nums ...int) {
	sum := 0

	for i := 0; i < len(nums); i++ {
		fmt.Println(nums[i])
		//sum = sum + nums[i]
		sum += nums[i]
	}

	fmt.Println("sum:", sum)
}
