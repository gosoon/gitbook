package main

func shortestSubarray(nums []int, k int) int {

	min := len(nums) + 1
	for i := 0; i < len(nums); i++ {
		//sum := nums[i]
		var sum int
		for j := i; j < len(nums); j++ {
			sum += nums[j]
			if sum >= k && j-i+1 < min {
				min = j - i + 1
			}
		}
	}

	if min == len(nums)+1 {
		return -1
	}
	return min

}
