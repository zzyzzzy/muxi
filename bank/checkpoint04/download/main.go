package main

import (
	"fmt"

	"github.com/weiji6/hacker-support/httptool"
)

func main() {

	fmt.Println("开始下载图片...")

	req, err := httptool.NewRequest(
		httptool.GETMETHOD,
		"http://http-theft-bank.gtainccnu.muxixyz.com/api/v1/organization/iris_sample",
		"",
		httptool.DEFAULT, // 这里可能不是 DEFAULT，自己去翻阅文档
	)
	if err != nil {
		fmt.Println(err)
	}
	req.SetHeader("passport", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJjb2RlIjoiZ2FvZmVuZ2ppYW5nIiwiaWF0IjoxNzYzMDA5MTc2LCJuYmYiOjE3NjMwMDkxNzZ9.YP18f01ah_WMB49tTeFQ5yOQJxSBdrNwl8blfNaideI")

	// write your code below
	// ...
	resp, err := req.SendRequest()
	if err != nil {

	}
	err = resp.Save("C:/Users/zhangziyue/Pictures/test.jpg") // 换成你自己的路径
	if err != nil {
		fmt.Println(err)
	}
}
