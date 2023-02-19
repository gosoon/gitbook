package main

/*
(preSumNums[q]−preSumNums[p])%k == 0
				|
				v
	preSumNums[q]%k == preSumNums[p]%k
*/
func checkSubarraySum(nums []int, k int) bool {
	preSumNums := make([]int, len(nums)+1)
	for i := 1; i <= len(nums); i++ {
		preSumNums[i] = preSumNums[i-1] + nums[i-1]
	}

	m := make(map[int]int)
	for i := 0; i < len(preSumNums); i++ {
		if index, found := m[preSumNums[i]%k]; found {
			if i-index >= 2 {
				return true
			}
		} else {
			m[preSumNums[i]%k] = i
		}
	}

	return false
}
