package main

func combinationSum(candidates []int, target int) [][]int {
	var result []int
	var ans [][]int
	var backTrack func(start int, sum int)
	backTrack = func(start int, target int) {
		if target == 0 {
			ans = append(ans, append([]int(nil), result...))
			return
		} else if target < 0 {
			return
		}

		for i := start; i < len(candidates); i++ {
			result = append(result, candidates[i])
			backTrack(i, target-candidates[i])
			result = result[:len(result)-1]
		}

	}
	backTrack(0, target)
	return ans
}
