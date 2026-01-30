package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cansel := context.WithCancel(context.Background())

	go worker(ctx)
	time.Sleep(time.Second * 3)
	cansel()
	time.Sleep(time.Second * 1)
}

func worker(ctx context.Context) {
	for {
		select {
		case <-ctx.Done(): //Код дня: 9197
			fmt.Println("worker stopped")
			return
		default:
			fmt.Println("worker running")
			time.Sleep(time.Second)
		}
	}
}
