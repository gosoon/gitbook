package main

/*
func subsets(nums []int) [][]int {
	preSets := []int{}
	ans := [][]int{}
	var backtrack func(start int)
	backtrack = func(start int) {
		ans = append(ans, append([]int{}, preSets...))

		for i := start; i < len(nums); i++ {
			preSets = append(preSets, nums[i])
			backtrack(i + 1)
			preSets = preSets[:len(preSets)-1]
		}
	}
	backtrack(0)
	return ans
}
*/
