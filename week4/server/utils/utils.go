// 这是工具包
package utils

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"muxi/common/message"
	"net"
)

// 这里将这些方法关联到结构体中
type Transfer struct {
	//分析应该有那些字段
	Conn net.Conn
	Buf  [8096]byte //这是传输时使用的缓冲

}

func (this *Transfer) ReadPkg() (mes message.Message, err error) {
	//buf := make([]byte, 8096)
	fmt.Println("读客户端发送的数据")
	_, err = this.Conn.Read(this.Buf[:4])
	if err != nil {
		//err = errors.New("read pkg header error")
		return
	}
	//根据buf[:4]转成一个uint32类型
	var pkgLen uint32 = binary.BigEndian.Uint32(this.Buf[0:4])
	//根据pkgLen读取消息内容
	n, err := this.Conn.Read(this.Buf[:pkgLen])
	if n != int(pkgLen) || err != nil {
		//err = errors.New("resd pkg body error")
		return
	}
	//把pkgLen反序列化成messag.Message
	err = json.Unmarshal(this.Buf[:pkgLen], &mes) //这个&非常重要
	if err != nil {
		fmt.Println("json.Unmarshal fail", err)
		return
	}
	return
}
func (this *Transfer) WritePkg(data []byte) (err error) {
	//先发送一个长度给对方
	var pkgLen uint32 = uint32(len(data))
	//var buf [4]byte
	binary.BigEndian.PutUint32(this.Buf[:4], pkgLen)
	//发送长度
	n, err := this.Conn.Write(this.Buf[:4])
	if n != 4 || err != nil {
		fmt.Println("conn.Write(bytes)fail", err)
		return
	}
	//发送data数据本身
	n, err = this.Conn.Write(data)
	if n != int(pkgLen) || err != nil {
		fmt.Println("conn.Write(bytes)fail", err)
		return
	}
	return
}
