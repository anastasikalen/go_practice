package main

import "fmt"

func main() {
	s := []string{"go", "python", "ruby"}
	n := []int{1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println(Filter(s, func(s string) bool {
		return len(s) > 3
	}))
	fmt.Println(Filter(n, func(n int) bool {
		return n%2 == 0
	}))
}
func Filter[T any](items []T, f func(T) bool) []T {
	sl := make([]T, 0, len(items))
	for _, item := range items {
		if f(item) {
			sl = append(sl, item)
		}
	}
	return sl
}
