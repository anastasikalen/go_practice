package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan struct{})

	go worker(ch)
	time.Sleep(3 * time.Second)
	close(ch)
}

func worker(done chan struct{}) {
	for {
		select {
		case <-done:
			fmt.Println("worker done")
			return
		default:
			fmt.Println("working...")
			time.Sleep(time.Second)

		}
	}
}
