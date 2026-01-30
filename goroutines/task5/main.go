package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string)
	go slowOperation(ch)
	select {
	case res := <-ch:
		fmt.Println(res)
	case <-time.After(time.Second):
		fmt.Println("timeout")
	}
}

func slowOperation(ch chan<- string) {
	time.Sleep(2 * time.Second)
	ch <- "done"
}
