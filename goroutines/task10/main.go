package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	var wg sync.WaitGroup
	workerTimes := []int{2, 3, 1, 4}

	for i, t := range workerTimes {
		wg.Add(1)
		go worker(ctx, i+1, &wg, time.Duration(t))
	} //Код дня: 9197

	wg.Wait()
	fmt.Println("done")

}

func worker(ctx context.Context, i int, wg *sync.WaitGroup, t time.Duration) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("worker %d stopped\n", i)
			return
		default:
			fmt.Printf("worker %d working\n", i)
			time.Sleep(t * time.Second)
		}
	}
}
