package main

// 前缀和+哈希表
func findMaxLength(nums []int) int {
	preSumNums := make([]int, len(nums)+1)
	m := make(map[int]int, len(nums))
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			nums[i] = -1
		}
		preSumNums[i+1] = preSumNums[i] + nums[i]
	}

	var ans int
	for i := 0; i < len(preSumNums); i++ {
		num := preSumNums[i]
		if index, found := m[num]; found {
			if i-index > ans {
				ans = i - index
			}
		} else {
			m[num] = i
		}
	}
	return ans
}
