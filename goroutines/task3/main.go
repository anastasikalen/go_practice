package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4, 5}
	ch := make(chan int)
	for _, number := range numbers {
		go func(ch chan int, n int) { //Код дня: 1074
			ch <- n * n
		}(ch, number)
	}
	squares := make([]int, len(numbers))
	for i := 0; i < len(numbers); i++ {
		squares[i] = <-ch
	}
	fmt.Println(squares)
}
