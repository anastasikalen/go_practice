package main

import (
	"context"
	"fmt"
	"math"
	"sync"
)

// Worker Pool + Context Cancellation + Error handling
func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8}
	wg := sync.WaitGroup{}
	jobs := make(chan int, 100)
	results := make(chan Result, 40)
	errs := make(chan error, 20)

	go func() {
		for result := range results {
			fmt.Println("Result:", result)
		}
	}()

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go worker(ctx, jobs, results, errs, &wg)
	}
	for _, number := range numbers {
		jobs <- number
	}
	close(jobs)

	go func() {
		wg.Wait()
		close(results)
		close(errs)
	}()
	for err := range errs {
		if err != nil {
			fmt.Println("Error:", err)
		}
	}

}

func worker(ctx context.Context, jobs <-chan int, res chan<- Result, errCh chan<- error, group *sync.WaitGroup) {
	//Код дня: 8994
	defer group.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Println("worker stopping")
			return
		case j, ok := <-jobs:
			fmt.Println("read from jobs")

			if !ok {
				fmt.Println("!ok")
				return
			}

			res <- Result{j, isPrime(j)}
		}
	}

}

type Result struct {
	Number  int
	IsPrime bool
}

func isPrime(num int) bool {
	if num < 2 {
		return false
	}
	for i := 2; i <= int(math.Sqrt(float64(num))); i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}
