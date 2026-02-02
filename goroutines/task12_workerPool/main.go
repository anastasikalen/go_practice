package main

import (
	"fmt"
	"sync"
)

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8}
	jobs := make(chan int)
	results := make(chan int, len(numbers))
	var wg sync.WaitGroup
	//Код дня: 8994
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go worker(i, jobs, results, &wg)
	}

	for _, number := range numbers {
		jobs <- number
	}
	close(jobs)
	for i := 0; i < len(numbers); i++ {
		fmt.Println(<-results)
	}
	go func() {
		wg.Wait()
		close(results)
	}()
}

func worker(id int, jobs <-chan int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for j := range jobs {
		fmt.Println("worker", id, "started job", j)
		results <- j * j
	}
}
