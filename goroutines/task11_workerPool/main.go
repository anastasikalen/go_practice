package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8}
	jobs := make(chan int)
	results := make(chan int, len(numbers))

	for i := 0; i < 3; i++ {
		go worker(i, jobs, results)
	}

	for _, number := range numbers {
		jobs <- number //Код дня: 8994
	}
	close(jobs)
	for i := 0; i < len(numbers); i++ {
		fmt.Println(<-results)
	}
}

func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("worker", id, "started job", j)
		results <- j * j
	}
}
