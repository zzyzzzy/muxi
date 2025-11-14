package main

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/Grand-Theft-Auto-In-CCNU-MUXI/hacker-support/encrypt"
)

func Put(secretkey string, errorCode string) {

	//加密
	data01, err := encrypt.AESEncryptOutInBase64([]byte(errorCode), []byte(secretkey))

	//data := encrypt.Base64Encode(data01)
	if err != nil {

		fmt.Println("加密失败", err)
	}
	fmt.Printf("加密后的 error_code: %s\n", data01)

	type Error_Code struct {
		Content string `json:"content"`
	}
	error_code := Error_Code{
		Content: string(data01),
	}
	//json
	jsonDate, err := json.Marshal(error_code)
	if err != nil {
		fmt.Println("序列化失败：", err)
	}
	url := "http://http-theft-bank.gtainccnu.muxixyz.com/api/v1/bank/gate"
	r, err := http.NewRequest("PUT", url, bytes.NewBuffer(jsonDate))

	if err != nil {
		fmt.Println("put失败：", err)
		return
	}
	client := &http.Client{}
	r.Header.Set("code", "助我破鼎")
	r.Header.Set("Content-Type", "application/json")
	r.Header.Set("passport", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJjb2RlIjoiMTHpmr7pgZPmmK_kvaDlrrPkuobmiJEiLCJpYXQiOjE3NjI5MTg4NjksIm5iZiI6MTc2MjkxODg2OX0.koIQNVwXpMRFESk_aK1PtYs9RKcBFUXyY7SApsgLPPM")
	res, err := client.Do(r)
	if err != nil {
		fmt.Println("发送请求失败：", err)
		return
	}
	defer res.Body.Close()
	resp, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("读body失败：", err)
	}
	fmt.Println("body:", string(resp))
	for key, values := range res.Header {
		fmt.Printf("  %s: %v\n", key, values)
	}
}
func main() {
	extraInfo := "c2VjcmV0X2tleTpNdXhpU3R1ZGlvMjAzMzA0LCBlcnJvcl9jb2RlOmZvciB7Z28gZnVuYygpe3RpbWUuU2xlZXAoMSp0aW1lLkhvdXIpfSgpfQ=="
	// Base64 解码
	decoded, err := base64.StdEncoding.DecodeString(extraInfo)

	if err != nil {
		fmt.Println("解码错误：", err)
		return
	}
	// 打印解码后的内容
	fmt.Println("解码后的内容:", string(decoded))
	//直接放这了
	secretkey := "MuxiStudio203304"
	errorCode := "for {go func(){time.Sleep(1*time.Hour)}()}"

	Put(secretkey, errorCode)

}
