package main

import "fmt"

func main() {
	nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	fmt.Println(removeDuplicates(nums))
}

func removeDuplicates(nums []int) int {
	for i := 0; i < len(nums); i++ {
		for j := len(nums) - 1; j >= 0; j-- {
			if nums[i] == nums[j] && i != j {
				nums = removeIdx(nums, j)
			}
		}
	}
	return len(nums)
}

func removeIdx(nums []int, i int) []int {
	nums = append(nums[:i], nums[i+1:]...)
	return nums
}
