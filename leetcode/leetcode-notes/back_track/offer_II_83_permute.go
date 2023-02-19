package main

import "fmt"

func permute(nums []int) [][]int {
	var ans [][]int
	var result []int
	used := make(map[int]bool)
	var backtrack func()
	backtrack = func() {
		fmt.Printf("result:%v \n", result)
		if len(nums) == len(result) {
			ans = append(ans, append([]int(nil), result...))
			return
		}

		for i := 0; i < len(nums); i++ {
			if !used[nums[i]] {
				result = append(result, nums[i])
				used[nums[i]] = true
				backtrack()
				result = result[:len(result)-1]
				used[nums[i]] = false
			}
		}
	}
	backtrack()
	return ans
}
