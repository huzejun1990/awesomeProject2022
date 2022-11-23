package main

import "fmt"

func main() {

	const (
		a = iota   // 0
		b          // 1
		c          // 2
		d = "hehe" // hehe	iota 3
		e          // hehe	iota 4
		f = 100    // 100	iota 5
		g          // 100  iota 6
		h = iota   // iota	7
		i          // iota  8
	)

	const (
		j = iota
		k
	)

	fmt.Println(a, b, c, d, e, f, g, h, i, j, k)

}
