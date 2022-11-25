package main

import "fmt"

func main() {

	//验证密码，再次输入密码
	var a, b int
	var pwd int = 20221124
	// 用户的输入
	fmt.Print("请输入密码：")
	fmt.Scan(&a)

	// 业务：验证密码是否正确
	if a == pwd {
		fmt.Print("请再次输入密码: ")
		fmt.Scan(&b)
		if b == pwd {
			fmt.Println("登陆成功了")
		} else {
			fmt.Println("登陆失败了，第二次密码错误")
		}
	} else {
		fmt.Println("登录失败，密码错误")
	}

}
