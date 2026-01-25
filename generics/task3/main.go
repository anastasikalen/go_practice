package main

import "fmt"

func main() {
	nums1 := []int{1, 2, 3}
	nums2 := []int{4, 5}
	fmt.Println(Concat(nums1, nums2))

	words1 := []string{"go", "java"}
	words2 := []string{"python", "ruby"}
	fmt.Println(Concat(words1, words2))

}
func Concat[T any](a, b []T) []T {
	sl := make([]T, len(a)+len(b))
	copy(sl, a)
	copy(sl[len(a):], b)
	return sl
}
