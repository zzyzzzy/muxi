package main

import (
	"fmt"
	"net"
)

func main() {
	//"tcp"协议
	//“127.0.0.1：8080”意思是我要去连127.0.0.1：8080这个服务器端口监听

	conn, err := net.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		fmt.Println("client err=", err)
		return
	}
	fmt.Println("conn成功", conn)
}
