package main

import (
	"fmt"
)

var userId int
var userPwd string

func main() {

	//接收用户选择
	var key int
	//判断是否还在继续显示菜单
	var loop = true
	for loop {
		fmt.Println("欢迎")
		fmt.Println("/t 1 登录聊天室")
		fmt.Println("/t 2 没有账号？注册用户")
		fmt.Println("/t 3 退出系统")
		fmt.Println("/t 请选择（1-3）")
		fmt.Scanf("%d\n", &key)
		switch key {
		case 1:
			fmt.Println("登录聊天室")
			loop = false
		case 2:
			fmt.Println("注册用户")
			loop = false
		case 3:
			fmt.Println("退出系统")
			loop = false
		default:
			fmt.Println("你的输入有误，请重新输入")

		}
	}
	if key == 1 {
		//登陆
		fmt.Println("请输入id：")
		fmt.Scanf("%d\n", &userId)
		fmt.Println("请输入密码：")
		fmt.Scanf("%s\n", &userPwd)
		//先把登录的函数写在另一个文件，login.go
		Login(userId, userPwd)
		// if err != nil {
		// 	fmt.Println("登录失败")
		// } else {
		// 	fmt.Println("登录成功")
		// }
	} else if key == 2 {
		fmt.Println("进行用户注册逻辑...")
	}
}
