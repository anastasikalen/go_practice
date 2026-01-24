package main

import "fmt"

func main() {

	sl := []int{1, 2, 2, 3, 1, 4}
	fmt.Println(Unique(sl))
}
func Unique(nums []int) []int { //выводить в новом слайсе только уникальные значения
	res := make([]int, 0, len(nums))
	m := map[int]bool{}

	for _, n := range nums {
		if m[n] {
			continue
		} else {
			m[n] = true
			res = append(res, n)
		}
	}
	return res
}
