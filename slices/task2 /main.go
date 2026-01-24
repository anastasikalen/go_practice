package main

import "fmt"

func main() {
	nums := []int{10, 20, 30, 40}
	fmt.Println(RemoveAt(nums, 1))
}
func RemoveAt(nums []int, index int) []int { //работа с индексами
	if index >= len(nums) || index < 0 {
		return nums
	}
	res := make([]int, 0, len(nums)-1)
	res = append(res, nums[:index]...)
	res = append(res, nums[index+1:]...)
	return res
}
