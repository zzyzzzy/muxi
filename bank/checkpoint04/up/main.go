package main

import (
	"fmt"

	"github.com/weiji6/hacker-support/httptool"
)

func main() {
	req, err := httptool.NewRequest(
		httptool.PUTMETHOD,
		"http://http-theft-bank.gtainccnu.muxixyz.com/api/v1/bank/iris_recognition_gate",
		"/Users/zhangziyue/Pictures/test.jpg",
		httptool.FILE)
	if err != nil {
		fmt.Println(err)
	}
	req.AddHeader("passport", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJjb2RlIjoiMTHpmr7pgZPmmK_kvaDlrrPkuobmiJEiLCJpYXQiOjE3NjI5MTg4NjksIm5iZiI6MTc2MjkxODg2OX0.koIQNVwXpMRFESk_aK1PtYs9RKcBFUXyY7SApsgLPPM")

	resp, err := req.SendRequest()
	if err != nil {
		fmt.Println(err)
	}
	// write your code below
	// ...

	resp.ShowBody()
	resp.ShowHeader()

}
