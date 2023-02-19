package main

func combinationSum4(nums []int, target int) int {
	m := make(map[int]int)
	var backtrack func(target int) int
	backtrack = func(target int) int {

		if target <= 0 {
			if target == 0 {
				return 1
			}
			return -1
		}

		if count, found := m[target]; found {
			return count
		}

		res := 0
		for i := 0; i < len(nums); i++ {
			result := backtrack(target - nums[i])
			if result != -1 {
				res += result
			}
		}
		m[target] = res
		return res
	}
	res := backtrack(target)
	if res == -1 {
		return 0
	}

	return res
}
