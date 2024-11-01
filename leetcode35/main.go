package main

import (
	"fmt"
	"slices"
)

func main() {
	var nums = []int{1, 3, 5, 6}
	var target int = 2

	res := searchInsert(nums, target)
	fmt.Println(res)
}

func searchInsert(nums []int, target int) int {
	nums = append(nums, target)
	slices.Sort(nums)
	res := 0
	for i, num := range nums {
		if num == target {
			res = i
			break
		}
	}
	return res
}
