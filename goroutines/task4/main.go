package main

import (
	"fmt"
	"sync"
)

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8}
	var wg sync.WaitGroup

	jobs := make(chan int, len(numbers))
	results := make(chan int)

	wg.Add(2)
	go worker(1, jobs, results, &wg)
	go worker(2, jobs, results, &wg)

	for _, number := range numbers {
		jobs <- number
	}
	close(jobs)

	go func() {
		wg.Wait()
		close(results)
	}()
	for result := range results {
		fmt.Println(result)
	}
}

func worker(id int, jobs <-chan int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for n := range jobs {
		results <- n * n
		fmt.Printf("worker %d processing %d\n", id, n)
	}
	close(results)
}
