package main

// Input: nums = [1,2,3,1]
// Output: true

func containsDuplicate(nums []int) bool {
	for i := 0; i < len(nums)-1; i++ {
		for j := len(nums) - 1; j > i; j-- {
			if nums[i] == nums[j] {
				return true
			}
		}
	}

	return false
}
