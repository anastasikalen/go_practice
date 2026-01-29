package main

import (
	"fmt"
	"sync"
)

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	ch := make(chan int)
	var wg sync.WaitGroup
	wg.Add(2)

	go func(ch chan int, nums []int) {
		defer wg.Done()
		sum := 0
		for _, number := range nums {
			sum += number
		}
		ch <- sum
	}(ch, numbers[:5])

	go func(nums []int) {
		defer wg.Done() //Код дня: 1074
		sum := 0
		for _, number := range nums {
			sum += number
		}
		ch <- sum

	}(numbers[5:])

	sum1 := <-ch
	sum2 := <-ch
	total := sum1 + sum2
	fmt.Println(total)
	wg.Wait()

}
