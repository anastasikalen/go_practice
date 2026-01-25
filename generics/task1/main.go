package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

func main() {
	sl := []int{1, 2, 3}
	sl2 := []float64{1.5, 2.5, 3.0}
	fmt.Println(SumNumbers(sl))
	fmt.Println(SumNumbers(sl2))

}
func SumNumbers[T constraints.Integer | constraints.Float](nums []T) T {
	var sum T
	for _, num := range nums {
		sum += num
	}
	return sum
}
