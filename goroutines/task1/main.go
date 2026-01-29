package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	ch := make(chan int)
	go func(ch chan int, numbers []int) {
		sum := 0
		for _, number := range numbers {
			sum += number
		}
		ch <- sum
	}(ch, numbers[:5])
	sum1 := <-ch //Код дня: 1074

	go func(ch chan int, numbers []int) {
		sum := 0
		for _, number := range numbers {
			sum += number
		}
		ch <- sum
	}(ch, numbers[5:])
	sum2 := <-ch
	total := sum1 + sum2
	fmt.Println(total)

}
