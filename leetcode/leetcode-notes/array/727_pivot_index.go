package main

import "fmt"

func pivotIndex(nums []int) int {
	preSum := make([]int, len(nums)+1)
	for i := 1; i <= len(nums); i++ {
		preSum[i] = nums[i-1] + preSum[i-1]
	}
	fmt.Println(preSum)
	for i := 0; i < len(nums); i++ {
		var leftSum, rightSum int
		if i == 0 {
			leftSum = preSum[i+1]
			rightSum = preSum[len(preSum)-1]
		} else {
			leftSum = preSum[i]
			rightSum = preSum[len(preSum)-1] - leftSum - nums[i]
		}
		fmt.Println(i, leftSum, rightSum)
		if leftSum == rightSum {
			return i
		}
	}
	return -1

}
