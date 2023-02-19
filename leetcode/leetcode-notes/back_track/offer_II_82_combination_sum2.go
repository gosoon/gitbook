package main

import "sort"

func combinationSum2(candidates []int, target int) [][]int {
	var ans [][]int
	var result []int
	sort.Ints(candidates)
	//numCount := make(map[int]int)
	//for _, num := range candidates {
	//numCount[num]++
	//}
	//globalUsedCount := make(map[int]int)

	var backtrack func(int, int)
	backtrack = func(start int, target int) {
		if target == 0 {
			ans = append(ans, append([]int(nil), result...))
			return
		} else if target < 0 {
			return
		}

		used := make(map[int]bool)
		for i := start; i < len(candidates); i++ {
			num := candidates[i]
			//count := numCount[num]
			//usedCount := globalUsedCount[num]
			//usedCount++
			if !used[num] {
				result = append(result, num)
				used[num] = true
				//globalUsedCount[num]++
				backtrack(i+1, target-num)
				//globalUsedCount[num]--
				result = result[:len(result)-1]
			}
		}
	}

	backtrack(0, target)
	return ans
}
