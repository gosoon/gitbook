package main

import "sort"

func twoSumUpgrade(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{}
	}
	sort.Ints(nums)

	l := 0
	r := len(nums) - 1

	for l < r {
		sum := nums[l] + nums[r]
		if sum < target {
			l++
		} else if sum > target {
			r--
		} else {
			return []int{nums[l], nums[r]}
		}
	}

	return []int{}
}
