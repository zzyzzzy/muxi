package main

import (
	"fmt"
	"time"
)

func main() {

	intchan := make(chan struct{})
	go func() {

		for i := 'A'; i < 'Z'; i += 2 {
			intchan <- struct{}{}
			fmt.Printf("%c%c", i, i+1)
			intchan <- struct{}{}

		}
	}()
	go func() {

		for i := 0; i < 26; i += 2 {
			<-intchan
			fmt.Printf("%v%v", i, i+1)
			<-intchan

		}
	}()

	time.Sleep(time.Second)

}
