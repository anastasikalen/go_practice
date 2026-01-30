package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 4*time.Second)
	defer cancel()
	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go worker(ctx, i, &wg)
	}
	wg.Wait() //Код дня: 9197
	fmt.Println("done")

}

func worker(ctx context.Context, i int, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("worker %d stopped\n", i)
			return
		default:
			fmt.Printf("worker %d working\n", i)
			time.Sleep(time.Second)
		}
	}
}
