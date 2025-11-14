package main

import (
	"github.com/weiji6/hacker-support/httptool"
)

func main() {
	r, err := httptool.NewRequest(
		httptool.PUTMETHOD,
		"http://http-theft-bank.gtainccnu.muxixyz.com/api/v1/muxi/backend/computer/examination",
		"C:/muxi/bank/checkpoint05/key/main.go",
		httptool.FILE,
	)

	if err != nil {
		panic(err)
	}
	r.SetHeader("passport", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJjb2RlIjoiZ2FvZmVuZ2ppYW5nIiwiaWF0IjoxNzYzMDA5MTc2LCJuYmYiOjE3NjMwMDkxNzZ9.YP18f01ah_WMB49tTeFQ5yOQJxSBdrNwl8blfNaideI")
	req, err := r.SendRequest()
	if err != nil {
		panic(err)
	}
	req.ShowBody()
	req.ShowHeader()
}
