package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func Get() {
	url := "http://http-theft-bank.gtainccnu.muxixyz.com/api/v1/organization/secret_key"
	r, err := http.NewRequest("PUT", url, nil)
	if err != nil {
		fmt.Println("get失败：", err)
		return
	}

	client := &http.Client{}
	r.Header.Set("code", "难道是你害了我")
	r.Header.Set("passport", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJjb2RlIjoiMTHpmr7pgZPmmK_kvaDlrrPkuobmiJEiLCJpYXQiOjE3NjI5MTg4NjksIm5iZiI6MTc2MjkxODg2OX0.koIQNVwXpMRFESk_aK1PtYs9RKcBFUXyY7SApsgLPPM")
	r.Header.Set("Content-Type", "application/json")
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
	var info map[string]interface{}
	//info=make(map[string]interface{})
	err = json.Unmarshal([]byte(resp), &info)
	if err != nil {
		fmt.Println("反序列化失败：", err)
	}
	fmt.Println("info=", info)

}

func main() {
	Get()
}
