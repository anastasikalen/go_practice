package main

import "fmt"

func main() {
	sl := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(FilterEven(sl))

}
func FilterEven(nums []int) []int { //вернуть новый слайс с четными числами
	sl := make([]int, 0)
	for _, n := range nums {
		if n%2 == 0 {
			sl = append(sl, n)
		}
	}
	return sl
}
