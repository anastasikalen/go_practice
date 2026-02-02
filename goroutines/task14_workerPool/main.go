package main

import (
	"context"
	"fmt"
	"sync"
)

// Worker Pool + Context Cancellation + Error handling
func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	jobs := make(chan int)
	errCh := make(chan error, len(nums))
	res := make(chan int, len(nums))
	var wg sync.WaitGroup
	//Код дня: 8994
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go worker(ctx, i, jobs, errCh, &wg, res)
	}

	for _, num := range nums {
		jobs <- num
	}
	close(jobs)

	go func() {
		wg.Wait()
		close(res)
		close(errCh)
	}()

	go func(errChan chan error) {
		for err := range errCh {
			fmt.Println(err.Error())
		}
	}(errCh)

	for r := range res {
		fmt.Println(r)
	}
}

func worker(ctx context.Context, id int, jobs <-chan int, errCh chan<- error, wg *sync.WaitGroup, res chan<- int) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Println("worker ", id, " stopping")
			return
		case j, ok := <-jobs:
			fmt.Println("read from jobs")

			if !ok {
				fmt.Println("!ok")
				return
			} else if j%5 == 0 {
				errCh <- fmt.Errorf("worker %d received job %d", id, j)
			} else {
				res <- j * j
			}
		}
	}
}
