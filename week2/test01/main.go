package main

import "fmt"

func tool() func(int) int {
	n := 0
	return func(x int) int {
		n += x
		return n
	}
}
func main() {
	f := tool()
	fmt.Println(f(1))
	fmt.Println(f(1))
}
