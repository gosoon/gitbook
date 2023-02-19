package main

import (
	"math"
	"sort"
)

/*
// three sum
func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)

	var res []int
	for i := 0; i < len(nums)-2; {
		resultList := twoSumFuncForApprochThreeSum(nums, i+1, target-nums[i])
		if len(res) == 0 {
			resultList = append(resultList, nums[i])
			res = resultList
		} else {
			var latestSum int
			for _, r := range res {
				latestSum += r
			}

			sum := nums[i]
			for _, result := range resultList {
				sum += result
			}

			if math.Abs(float64(sum)-float64(target)) < math.Abs(float64(latestSum)-float64(target)) {
				resultList = append(resultList, nums[i])
				res = resultList
			}
		}

		// skip repeat
		left := nums[i]
		for i < len(nums)-1 && nums[i] == left {
			i++
		}
	}
	var sum int
	for _, r := range res {
		sum += r
	}
	return sum
}

// two sum
func twoSumFuncForApprochThreeSum(nums []int, start int, target int) []int {
	l := start
	r := len(nums) - 1

	var res []int
	for l < r {
		left := nums[l]
		right := nums[r]
		sum := left + right

		if len(res) == 0 {
			res = []int{left, right}
		} else {
			var latestSum int
			for _, r := range res {
				latestSum += r
			}

			if math.Abs(float64(sum)-float64(target)) < math.Abs(float64(latestSum)-float64(target)) {
				res = []int{left, right}
			}
		}

		if sum < target {
			l++
		} else if sum > target {
			r--
		} else {
			return res
		}
	}
	return res
}
*/

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	closest := nums[0] + nums[1] + nums[2]
	if closest >= target {
		return closest
	}

	for i := 0; i < len(nums)-2; i++ {
		l := i + 1
		r := len(nums) - 1

		for l < r {
			sum := nums[l] + nums[r] + nums[i]

			if math.Abs(float64(sum)-float64(target)) < math.Abs(float64(closest)-float64(target)) {
				closest = sum
			}

			if sum < target {
				l++
			} else if sum > target {
				r--
			} else {
				return sum
			}
		}
	}
	return closest
}
