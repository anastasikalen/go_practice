package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cansel := context.WithTimeout(context.Background(), 4*time.Second)
	defer cansel()
	for i := 0; i < 5; i++ {
		go worker(ctx, i)
	}
	context.WithTimeout(ctx, time.Second*4)
	time.Sleep(time.Second * 5)

	fmt.Println("done")
}

// Код дня: 9197
func worker(ctx context.Context, i int) {
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("worker %d stopped\n", i)
			return
		default:
			fmt.Printf("worker %d running\n", i)
			time.Sleep(time.Second)
		}
	}
}
