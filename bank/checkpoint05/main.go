package main

import (
	//"fmt"
	"github.com/weiji6/hacker-support/httptool"
)

func main() {
	request, err := httptool.NewRequest(
		httptool.GETMETHOD,
		"http://http-theft-bank.gtainccnu.muxixyz.com/api/v1/muxi/backend/computer/examination",
		"",
		httptool.DEFAULT)
	if err != nil {
		panic(err)
	}
	request.SetHeader("passport", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJjb2RlIjoiemhhbmd6aXl1ZSIsImlhdCI6MTc2MzEwNTE4MCwibmJmIjoxNzYzMTA1MTgwfQ.jwodHBmM_1twP_ra_eXPh1vY7xocqvjZitnN5wbKHUU")
	resp, err := request.SendRequest()
	if err != nil {
		panic(err)
	}
	resp.ShowBody()
	resp.ShowHeader()
}
