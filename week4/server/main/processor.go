// 这里是主控
package main

import (
	"fmt"
	"io"
	"muxi/common/message"
	"muxi/week4/server/process"
	"muxi/week4/server/utils"
	"net"
)

// 先创建一个Processor的结构体
type Processor struct {
	Conn net.Conn
}

// 编写一个ServerProcessMes函数
// 功能是根据客户端发送的消息种类不同，调用不同的函数
func (this *Processor) ServerProcessMes(mes *message.Message) (err error) {
	switch mes.Type {
	case message.LoginMesType:
		//处理登录
		//创建一个UserProcess实例
		up := &process.UserProcess{
			Conn: this.Conn,
		}
		err = up.ServerProcessLogin(mes)
	case message.RegisterMesType:
		//处理注册
	default:
		fmt.Println("消息不存在，无法处理..")
	}
	return
}
func (this *Processor) process2() (err error) {
	//循环给客户端发送消息
	for {
		//这里我们将读取数据包，直接封装成一个函数readpkg(),返回Message，Err
		//创建一个Transfer 实例完成读包
		tf := &utils.Transfer{
			Conn: this.Conn,
		}
		mes, err := tf.ReadPkg()
		if err != nil {
			if err == io.EOF {
				fmt.Println("客户端退出，我也退出..")
				return err
			} else {
				fmt.Println("readPkg err", err)
				return err
			}
		}
		//fmt.Println("mes=", mes)
		err = this.ServerProcessMes(&mes)
		if err != nil {
			return err
		}
	}
}
