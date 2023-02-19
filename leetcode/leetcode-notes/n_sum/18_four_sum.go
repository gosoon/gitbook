package main

import (
	"sort"
)

// four sum
func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	var res [][]int
	for i := 0; i < len(nums)-1; {
		resultList := threeSumCommonFunc(nums, i+1, target-nums[i])
		for _, result := range resultList {
			result = append([]int{nums[i]}, result...)
			res = append(res, result)
		}

		// skip repeat
		left := nums[i]
		for i < len(nums)-1 && nums[i] == left {
			i++
		}
	}
	return res
}

// three sum
func threeSumCommonFunc(nums []int, start int, target int) [][]int {
	var res [][]int
	for i := start; i < len(nums)-1; {
		resultList := twoSumCommonFunc(nums, i+1, target-nums[i])

		for _, result := range resultList {
			result = append([]int{nums[i]}, result...)
			res = append(res, result)
		}
		// skip repeat
		left := nums[i]
		for i < len(nums)-1 && nums[i] == left {
			i++
		}
	}
	return res
}

// two sum
func twoSumCommonFunc(nums []int, start int, target int) [][]int {
	//sort.Ints(nums)
	l := start
	r := len(nums) - 1

	var res [][]int
	for l < r {
		left := nums[l]
		right := nums[r]
		sum := left + right
		if sum < target {
			l++
			// skip repeat nums
			for l < r && nums[l] == left {
				l++
			}
		} else if sum > target {
			r--
			// skip repeat nums
			for l < r && nums[r] == right {
				r--
			}
		} else {
			res = append(res, []int{left, right})
			// skip repeat nums
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
