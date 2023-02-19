package main

import "fmt"

func containsNearbyAlmostDuplicate(nums []int, indexDiff int, valueDiff int) bool {
	fmt.Printf("nums:%v,indexDiff:%v,valueDiff:%v \n", nums, indexDiff, valueDiff)

	for i := 0; i < len(nums); i++ {
		for j := i + 1; j <= getMin(i+indexDiff, len(nums)-1); j++ {
			if abs(nums[j]-nums[i]) <= valueDiff {
				return true
			}
		}
	}
	return false
}

func getMin(i, j int) int {
	if i < j {
		return i
	}
	return j
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}
