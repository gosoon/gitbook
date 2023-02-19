package main

func combine(n int, k int) [][]int {
	var results [][]int
	var preSet []int
	var backtrack func(start int)
	backtrack = func(start int) {
		if len(preSet) == k {
			results = append(results, append([]int(nil), preSet...))
			return
		}

		for i := start; i <= n; i++ {
			preSet = append(preSet, i)
			backtrack(i + 1)
			preSet = preSet[:len(preSet)-1]
		}
	}
	backtrack(1)
	return results
}
