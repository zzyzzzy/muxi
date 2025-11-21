package main

import (
	//"encoding/binary"
	//"encoding/json"

	//"errors"
	"fmt"
	//"io"
	//"muxi/common/message"
	"net"
	//"golang.org/x/text/cases"
)

// var peo01 = message.LoginMes{
// 	UserId:   100,
// 	UserPwd:  "123456",
// 	UserName: "小张",
// }

// 把读取消息封装成一个函数
// func readPkg(conn net.Conn) (mes message.Message, err error) {
// 	buf := make([]byte, 8096)
// 	fmt.Println("读客户端发送的数据")
// 	_, err = conn.Read(buf[:4])
// 	if err != nil {
// 		//err = errors.New("read pkg header error")
// 		return
// 	}
// 	//根据buf[:4]转成一个uint32类型
// 	var pkgLen uint32 = binary.BigEndian.Uint32(buf[0:4])
// 	//根据pkgLen读取消息内容
// 	n, err := conn.Read(buf[:pkgLen])
// 	if n != int(pkgLen) || err != nil {
// 		//err = errors.New("resd pkg body error")
// 		return
// 	}
// 	//把pkgLen反序列化成messag.Message
// 	err = json.Unmarshal(buf[:pkgLen], &mes) //这个&非常重要
// 	if err != nil {
// 		fmt.Println("json.Unmarshal fail", err)
// 		return
// 	}
// 	return
// }
// func writePkg(conn net.Conn, data []byte) (err error) {
// 	//先发送一个长度给对方
// 	var pkgLen uint32 = uint32(len(data))
// 	var buf [4]byte
// 	binary.BigEndian.PutUint32(buf[:4], pkgLen)
// 	//发送长度
// 	n, err := conn.Write(buf[:4])
// 	if n != 4 || err != nil {
// 		fmt.Println("conn.Write(bytes)fail", err)
// 		return
// 	}
// 	//发送data数据本身
// 	n, err = conn.Write(data)
// 	if n != int(pkgLen) || err != nil {
// 		fmt.Println("conn.Write(bytes)fail", err)
// 		return
// 	}
// 	return
// }

// // 编写一个serverProcessLogin函数，专门处理登录请求
// func serverProcessLogin(conn net.Conn, mes *message.Message) (err error) {
// 	//核心代码
// 	//1.先从mes中取出mes.Data，并直接反序列化成LoginMes
// 	var loginMes message.LoginMes
// 	err = json.Unmarshal([]byte(mes.Data), &loginMes)
// 	if err != nil {
// 		fmt.Println("json.Unmarshal fail err=", err)
// 		return
// 	}
// 	//先声明一个resMes
// 	var resMes message.Message
// 	resMes.Type = message.LoginResMesType
// 	//再声明一个loginResMes
// 	var loginResMes message.LoginResMes

// 	//嘻嘻小张登场！！！！
// 	//仅为试验
// 	if loginMes.UserId == 100 && loginMes.UserPwd == "123456" {
// 		//合法
// 		loginResMes.Code = 200

// 	} else {
// 		//不合法
// 		loginResMes.Code = 500 //500表示该用户不存在
// 		loginResMes.Error = "该用户不存在，请注册再使用"
// 	}
// 	//3.将loginResMes序列化
// 	data, err := json.Marshal(loginResMes)
// 	if err != nil {
// 		fmt.Println("json.Marshal fail", err)
// 		return
// 	}
// 	//4.将data赋值给resMes
// 	resMes.Data = string(data)
// 	//5.对resMes进行序列化，准备发送
// 	data, err = json.Marshal(resMes)
// 	if err != nil {
// 		fmt.Println("json.Marshal fail", err)
// 		return
// 	}
// 	//6.发送
// 	err = writePkg(conn, data)
// 	return
// }

// // 编写一个ServerProcessMes函数
// // 功能是根据客户端发送的消息种类不同，调用不同的函数
// func ServerProcessMes(conn net.Conn, mes *message.Message) (err error) {
// 	switch mes.Type {
// 	case message.LoginMesType:
// 		//处理登录
// 	case message.RegisterMesType:
// 		//处理注册
// 	default:
// 		fmt.Println("消息不存在，无法处理..")
// 	}
// 	return
// }

// 处理和客户端通信
func processes(conn net.Conn) {
	//这里需要延时关闭
	defer conn.Close()
	//这里调用总控，创建一个
	processor := &Processor{
		Conn: conn,
	}
	err := processor.process2()
	if err != nil {
		fmt.Println("客户端和服务器通讯协程错误=err", err)
		return
	}


}

func main() {

	//提示信息
	fmt.Println("服务器[新的结构]在8889端口监听...")
	listen, err := net.Listen("tcp", "0.0.0.0:8889")
	//defer listen.Close()
	if err != nil {
		fmt.Println("net listen err=", err)
		return
	}
	//一旦监听成功，就等待客户端来连接服务器
	for {
		fmt.Println("等待客户端来连接服务器...")
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("listen.Accept err=", err)
		}
		//一旦连接成功则开启一个协程和客户端保持通讯
		go processes(conn)
	}
}
