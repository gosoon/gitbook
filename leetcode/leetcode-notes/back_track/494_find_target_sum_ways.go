package main

func findTargetSumWays(nums []int, target int) int {
	var result int
	var sum int
	findTargetSumWaysBackTrack(nums, 0, target, sum, &result)
	return result
}

func findTargetSumWaysBackTrack(nums []int, start int, target int, sum int, result *int) {
	// some return state
	if start == len(nums) {
		if sum == target {
			*result += 1
		}
		return
	}

	num := nums[start]
	findTargetSumWaysBackTrack(nums, start+1, target, sum+num, result)
	findTargetSumWaysBackTrack(nums, start+1, target, sum-num, result)
}
