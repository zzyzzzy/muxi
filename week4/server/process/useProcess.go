package process

//客户端请求
import (
	"encoding/json"
	"fmt"
	"muxi/common/message"
	"muxi/week4/server/utils"
	"net"
)

type UserProcess struct {
	//字段
	Conn net.Conn
}

// 编写一个serverProcessLogin函数，专门处理登录请求
func (this *UserProcess) ServerProcessLogin(mes *message.Message) (err error) {
	//核心代码
	//1.先从mes中取出mes.Data，并直接反序列化成LoginMes
	var loginMes message.LoginMes
	err = json.Unmarshal([]byte(mes.Data), &loginMes)
	if err != nil {
		fmt.Println("json.Unmarshal fail err=", err)
		return
	}
	//先声明一个resMes
	var resMes message.Message
	resMes.Type = message.LoginResMesType
	//再声明一个loginResMes
	var loginResMes message.LoginResMes

	//嘻嘻小张登场！！！！
	//仅为试验
	if loginMes.UserId == 100 && loginMes.UserPwd == "123456" {
		//合法
		loginResMes.Code = 200

	} else {
		//不合法
		loginResMes.Code = 500 //500表示该用户不存在
		loginResMes.Error = "该用户不存在，请注册再使用"
	}
	//3.将loginResMes序列化
	data, err := json.Marshal(loginResMes)
	if err != nil {
		fmt.Println("json.Marshal fail", err)
		return
	}
	//4.将data赋值给resMes
	resMes.Data = string(data)
	//5.对resMes进行序列化，准备发送
	data, err = json.Marshal(resMes)
	if err != nil {
		fmt.Println("json.Marshal fail", err)
		return
	}
	//6.发送,
	//因为使用分层模式（mvc），我们先创建一个Trankfer实例，然后读取
	tf := &utils.Transfer{
		Conn: this.Conn,
	}
	err = tf.WritePkg(data)
	return
}
