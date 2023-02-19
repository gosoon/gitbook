package main

func findSubsequences(nums []int) [][]int {
	findSubsequencesAns := [][]int{}
	var backTrack func(start int, subSet []int)
	backTrack = func(start int, subSet []int) {
		if len(subSet) > 1 {
			var result []int
			findSubsequencesAns = append(findSubsequencesAns, append(result, subSet...))
		}

		used := make(map[int]bool)
		for i := start; i < len(nums); i++ {
			if (len(subSet) > 0 && nums[i] < subSet[len(subSet)-1]) || used[nums[i]] {
				continue
			}
			backTrack(i+1, append(subSet, nums[i]))
			used[nums[i]] = true
		}
	}

	backTrack(0, []int{})
	return findSubsequencesAns
}
