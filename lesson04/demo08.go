package main

import "fmt"

func main() {

	sum := getSum2(5)
	fmt.Println(sum)

}

// 5
// 9 + 5
// 5 + 4
// 2 + 3
// 1 + 1
// 1
func getSum2(n int) int {
	if n == 1 {
		return 1
	}
	return getSum2(n-1) + n
}
