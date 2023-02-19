package main

/*
(preSumNums[q]−preSumNums[p])%k == 0
				|
				v
	preSumNums[q]%k == preSumNums[p]%k
*/

func subarraysDivByK(nums []int, k int) int {
	var availableSubArrays int

	var sum int
	m := map[int]int{0: 1}
	for _, num := range nums {
		sum += num
		//modulus := sum % k
		modulus := (sum%k + k) % k
		m[modulus]++
	}

	for _, count := range m {
		availableSubArrays += count * (count - 1) / 2
	}

	return availableSubArrays
}
