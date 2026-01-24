package main

import "fmt"

func main() {
	sl := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(GroupByParity(sl))
}
func GroupByParity(nums []int) map[string][]int { //возвращаем мапу с ключем четных и нечетных чисел из слайса
	m := make(map[string][]int)
	for _, n := range nums {
		if n%2 == 0 {
			m["even"] = append(m["even"], n)
		} else {
			m["odd"] = append(m["odd"], n)
		}
	}
	return m
}
