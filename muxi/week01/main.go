package main

import (
	"fmt"
)

func main() {
	// 给定s := make([]byte, 5)，那么len(s) 和 cap(s) 分别是多少？令s = s[2:4]，len(s) 和 cap(s) 又分别是多少？
	// 比较字符串"hello，世界"的长度和for range该字符串的循环次数
	s := make([]byte, 5)
	fmt.Println(len(s)) //5
	fmt.Println(cap(s)) //5

	s = s[2:4]
	fmt.Println(len(s)) //2
	fmt.Println(cap(s)) //3
	str := "hello,世界"

	fmt.Println("长度为:", len(str)) //12

	num := 0

	for range str {
		num++
	}
	fmt.Println("循环次数为:", num) //8

}
