package main

import (
	"sort"
)

// Input: nums = [9,6,4,2,3,5,7,0,1]
// Output: 8
// [1,2,4,5]

func missingNumber(nums []int) int {
	if len(nums) == 1 {
		if nums[0] == 1 {
			return 0
		}
		if nums[0] == 0 {
			return 1
		}
	}
	sort.Ints(nums)
	if nums[0] == 1 {
		return 0
	}
	for i := 0; i < len(nums)-1; i++ {
		if nums[i+1]-nums[i] != 1 {
			return nums[i] + 1
		}
	}
	return nums[len(nums)-1] + 1
}
