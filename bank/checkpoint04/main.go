package main

import (
	"fmt"
	"io"
	"net/http"
)

func Get() {
	url := "http://http-theft-bank.gtainccnu.muxixyz.com/api/v1/bank/iris_recognition_gate"
	r, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("get失败：", err)
		return
	}

	client := &http.Client{}
	r.Header.Set("code", "11难道是你害了我")
	r.Header.Set("passport", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJjb2RlIjoiMTHpmr7pgZPmmK_kvaDlrrPkuobmiJEiLCJpYXQiOjE3NjI5MTg4NjksIm5iZiI6MTc2MjkxODg2OX0.koIQNVwXpMRFESk_aK1PtYs9RKcBFUXyY7SApsgLPPM")
	res, err := client.Do(r)
	if err != nil {
		fmt.Println("发送请求失败：", err)
		return
	}
	defer res.Body.Close()
	resp, err := io.ReadAll(res.Body)
	// for key, values := range res.Header {
	// 	fmt.Printf("  %s: %v\n", key, t
	if err != nil {
		fmt.Println("读body失败：", err)
	}
	fmt.Println("body:", string(resp))
	for key, values := range res.Header {
		fmt.Printf("  %s: %v\n", key, values)
	}
}
func main() {
	Get()
}
