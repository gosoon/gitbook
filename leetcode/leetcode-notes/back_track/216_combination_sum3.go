package main

func combinationSum3(k int, n int) [][]int {
	var results [][]int
	var preSet []int
	var backtrack func(start int, sum int)
	backtrack = func(start int, sum int) {
		if len(preSet) == k {
			if sum == n {
				results = append(results, append([]int(nil), preSet...))
			}
			return
		}

		for i := start; i <= 9; i++ {
			preSet = append(preSet, i)
			backtrack(i+1, sum+i)
			preSet = preSet[:len(preSet)-1]
		}
	}
	backtrack(1, 0)
	return results
}
