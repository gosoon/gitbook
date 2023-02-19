package main

import (
	"sort"
)

func threeSum(nums []int) [][]int {
	var res [][]int
	sort.Ints(nums)
	r := len(nums)
	target := 0
	for l := 0; l < r; l++ {
		subTarget := target - nums[l]
		subResList := twoSumCommon(nums, l+1, subTarget)
		for _, subRes := range subResList {
			subRes = append([]int{nums[l]}, subRes...)
			res = append(res, subRes)
		}
		for l < r-1 && nums[l] == nums[l+1] {
			l++
		}
	}
	return res
}

func twoSumCommon(nums []int, start int, target int) [][]int {
	l := start
	r := len(nums) - 1

	var res [][]int
	for l < r {
		left := nums[l]
		right := nums[r]
		sum := left + right
		if sum < target {
			for l < r && nums[l] == left {
				l++
			}
		} else if sum > target {
			for l < r && nums[r] == right {
				r--
			}
		} else {
			res = append(res, []int{nums[l], nums[r]})
			for l < r && nums[l] == left {
				l++
			}
			for l < r && nums[r] == right {
				r--
			}
		}
	}
	return res
}
