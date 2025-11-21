package main

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"muxi/common/message"
	"net"
)

// 写一个函数完成登录
func Login(UesrId int, UserPwd string) (err error) {
	// //下一个就开始定协议
	// fmt.Println("UserId=%d UserPwd=%s\n", UserId, UserPwd)
	// return nil
	//1.连接服务器
	conn, err := net.Dial("tcp", "localhost:8889")
	if err != nil {
		fmt.Println("net.Dial err=", err)
		return
	}
	defer conn.Close()
	//2.准备通过conn发送消息给服务器
	var mes message.Message //message包包下面Message结构体
	mes.Type = message.LoginMesType
	//3.创建一个LoginMes结构体
	var loginMes message.LoginMes
	loginMes.UserId = UesrId
	loginMes.UserPwd = UserPwd
	//4.将loginmes序列化
	data, err := json.Marshal(loginMes)
	if err != nil {
		fmt.Println("json.Marshal err:", err)
		return
	}
	//5.把data赋给mes.Data字段
	mes.Data = string(data)
	//6.将mes.Data序列化
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("json.Marshal err:", err)
		return
	}
	//7.这个data就是我们要发送的消息
	//7.1先把data的长度发送给服务器
	//因为conn.Write要一个[]byte,所以现在我们要先获取data的长度，再转成一个表示长度的byte切片
	var pkgLen uint32 = uint32(len(data))
	var buf [4]byte
	binary.BigEndian.PutUint32(buf[:4], pkgLen)
	//发送长度
	n, err := conn.Write(buf[:4])
	if n != 4 || err != nil {
		fmt.Println("conn.Write(bytes)fail", err)
		return
	}
	fmt.Println("客户端发送的消息长度ok")
	//发送消息本身
	_, err = conn.Write(data)
	if err != nil {
		fmt.Println("conn.Write(bytes)fail", err)
		return
	}
	//这里还需要处理服务器返回的消息
	mes, err = readPkg(conn) //mes
	if err != nil {
		fmt.Println("resdPkg(conn)err=", err)
		return
	}
	//将mes的data部分反序列化成LoginResMes
	var loginResMes message.LoginResMes
	err = json.Unmarshal([]byte(mes.Data), &loginResMes)
	if loginResMes.Code == 200 {
		fmt.Println("登录成功")
	} else if loginResMes.Code == 500 {
		fmt.Println(loginResMes.Error)
	}
	return
}
