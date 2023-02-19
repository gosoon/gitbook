package main

import "math"

func maxSubArray(nums []int) int {
	dp := make([]int, len(nums))

	for i := 0; i < len(nums); i++ {
		dp[i] = nums[i]
	}

	for i := 1; i < len(nums); i++ {
		dp[i] = maxForMaxSubArray(dp[i], dp[i-1]+nums[i])
	}

	res := math.MinInt32
	for i := 0; i < len(dp); i++ {
		res = maxForMaxSubArray(res, dp[i])
	}
	return res
}

func maxForMaxSubArray(x, y int) int {
	if x > y {
		return x
	}
	return y
}
