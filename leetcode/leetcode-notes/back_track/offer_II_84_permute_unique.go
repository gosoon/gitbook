package main

import "fmt"

// 1.同层不能有重复值
// 2.全局每个值的使用数量不能超过 nums 中单个值的数量
func permuteUnique(nums []int) [][]int {
	var ans [][]int
	var result []int
	numCount := make(map[int]int)
	for _, num := range nums {
		numCount[num]++
	}
	globalUsed := make(map[int]int)

	var backtrack func()
	backtrack = func() {
		fmt.Printf("result:%v \n", result)
		if len(nums) == len(result) {
			ans = append(ans, append([]int(nil), result...))
			return
		}

		used := make(map[int]bool)
		for i := 0; i < len(nums); i++ {
			num := nums[i]
			usedCount := globalUsed[num]
			count := numCount[num]
			usedCount++
			if usedCount <= count && !used[num] {
				result = append(result, num)
				used[num] = true
				globalUsed[num]++
				backtrack()
				result = result[:len(result)-1]
				globalUsed[num]--
			}
		}
	}
	backtrack()
	return ans
}
