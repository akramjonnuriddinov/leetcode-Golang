package main

import "fmt"

func main() {
	nums := []int{1, 3, 4, 2}
	target := 6
	res := twoSum(nums, target)
	fmt.Println(res)
}

func twoSum(nums []int, target int) []int {
	result := []int{}
	for i := 0; i < len(nums); i++ {
		for j := len(nums) - 1; j > 0; j-- {
			if nums[i]+nums[j] == target && i != j {
				result = []int{i, j}
			}
		}
	}
	return result
}
