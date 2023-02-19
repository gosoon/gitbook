package main

import "fmt"

func unequalTriplets(nums []int) int {
	var ans int
	var backtrack func(index int, result []int)
	used := make(map[int]int)
	backtrack = func(index int, result []int) {
		if len(result) == 3 {
			fmt.Printf("result:%+v \n", result)
			ans++
			return
		}

		for i := index; i < len(nums); i++ {
			if _, found := used[nums[i]]; !found {
				used[nums[i]]++
				backtrack(i+1, append(result, nums[i]))
				delete(used, nums[i])
			}
		}
	}
	backtrack(0, []int(nil))
	return ans
}
