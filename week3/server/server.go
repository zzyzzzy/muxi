package main

import (
	"fmt"
	"net"
)

func main() {
	fmt.Println("服务器开始监听..")
	//"tcp"为协议，"127.0.0.1：8080"表示从本地监听8080端口
	listen, err := net.Listen("tcp", "127.0.0.1:8080")
	if err != nil {
		fmt.Println("listen err=", err)
		return //如果监听都出错，那就别玩了
	}
	defer listen.Close() //延迟关闭
	//循环等待客户端链接
	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("Accept() err=", err)
		} else {
			fmt.Printf("conn=%v", conn)
		}
	}
}
