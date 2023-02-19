package main

// ans 1
func subsets(nums []int) (ans [][]int) {
	var dfs func(i int)
	var tmp []int
	var res [][]int
	dfs = func(i int) {
		if i == len(nums) {
			res = append(res, append([]int(nil), tmp...))
			return
		}
		tmp = append(tmp, nums[i])
		dfs(i + 1)
		tmp = tmp[:len(tmp)-1]
		dfs(i + 1)
	}
	dfs(0)
	return res
}

/*
// ans 2
var ans [][]int

func subsets(nums []int) [][]int {
	preSets := []int{}
	ans = [][]int{}
	backtrack(nums, 0, preSets)
	return ans
}

func backtrack(nums []int, start int, preSets []int) {
	fmt.Printf("nums:%v start:%v preSets:%v \n", nums, start, preSets)
	ans = append(ans, append([]int{}, preSets...))

	for i := start; i < len(nums); i++ {
		//preSets = append(preSets, nums[i])
		backtrack(nums, i+1, append(preSets, nums[i]))
		//preSets = preSets[:len(preSets)-1]
	}
}
*/
