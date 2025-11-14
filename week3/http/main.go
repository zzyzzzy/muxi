package main

import (
	"fmt"
	"io"
	"net/http"
)

func Get() {
	r, err := http.Get("http://httpbin.org/get")
	if err != nil {
		panic(err)
		// fmt.Println("get err=",err)
	}
	defer r.Body.Close()
	content, err := io.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println("constent=", string(content))
}
func Post() {
	//倒数第二个双引号里面写post的东西的类型，比如jpg，json
	//最后一个东西是请求体数据，即HTTP请求中实际要发送给服务器的内容数据。
	r, err := http.Post("http://httpbin.org/post", "", nil)
	if err != nil {
		panic(err)
	}
	defer r.Body.Close()
	body, err := io.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println("body=", string(body))
}
func Put() {
	request, err := http.NewRequest(http.MethodPut, "http://httpbin.org/put&", nil)
	if err != nil {
		panic(err)
	}
	r, err := http.DefaultClient.Do(request)
	if err != nil {
		panic(err)
	}
	defer r.Body.Close()
	body, err := io.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println("body=", string(body))

}
func del() {
	//直接把put的复制粘贴过来，把methodput改成methoddelete就好
	request, err := http.NewRequest(http.MethodDelete, "http://httpbin.org/delete", nil)
	if err != nil {
		panic(err)
	}
	r, err := http.DefaultClient.Do(request)
	if err != nil {
		panic(err)
	}
	defer r.Body.Close()
	body, err := io.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println("body=", string(body))
}
func main() {
	Get()
	Put()
	Post()
	del()

}
