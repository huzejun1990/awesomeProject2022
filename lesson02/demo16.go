package main

import "fmt"

func main() {

	var a bool = false
	var b bool = true

	// 逻辑 && 与， 我与你，我和你都要满足一个结果，才执行
	// a 和 b 都为真的时候，结果为真，true,只要有一个为假，结果就为假
	// 都为真，结果才为真，否则就是假
	/*	if a == true && b == true {
		fmt.Println(a && b)
	}*/
	// 逻辑 || 或，我或者你，其中一个满足就可以
	// a || b,只要有一个为真，结果就为真，只有全是false,才会返回false
	//fmt.Println(a || b)

	// 逻辑非 ! !a 取反 如果为真，结果就为假，取反
	fmt.Println(!a)
	fmt.Println(!b)
}
