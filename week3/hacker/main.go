package main

import (
	"fmt"
	"io"
	"net/http"
)

func Get() {
	url := "https://gtainmuxi.muxixyz.com/api/v1/organization/code"
	r, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("get失败：", err)
		return
	}

	client := &http.Client{}
	r.Header.Set("code", "11难道是你害了我")
	r.Header.Set("passport", "11难道是你害了我")
	res, err := client.Do(r)
	if err != nil {
		fmt.Println("发送请求失败：", err)
		return
	}
	defer res.Body.Close()
	resp, err := io.ReadAll(res.Body)
	// for key, values := range res.Header {
	// 	fmt.Printf("  %s: %v\n", key, values)
	if err != nil {
		fmt.Println("读取body失败：", err)
		return
	}
	fmt.Println("Body:", string(resp))

}

func main() {
	Get()
}
