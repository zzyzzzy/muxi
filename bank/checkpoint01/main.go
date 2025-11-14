package main

import (
	//"fmt"
	"github.com/weiji6/hacker-support/httptool"
)

func main() {
	request, err := httptool.NewRequest(
		httptool.GETMETHOD,
		"https://gtainmuxi.muxixyz.com/api/v1/organization/code",
		"",
		httptool.DEFAULT)
	if err != nil {
		panic(err)
	}
	resp, err := request.SendRequest()
	if err != nil {
		panic(err)
	}
	resp.ShowBody()
	resp.ShowHeader()
}
