package main

import "fmt"

func canPartitionKSubsets(nums []int, k int) bool {
	if k > len(nums) {
		return false
	}

	var sum int
	for _, num := range nums {
		sum += num
	}
	if sum%k != 0 {
		return false
	}

	var backtrack func(index int) bool

	target := sum / k
	buckets := make([]int, k)
	backtrack = func(index int) bool {
		fmt.Printf("index:%v buckets:%+v \n", index, buckets)
		if index == len(nums) {
			for _, bucket := range buckets {
				if bucket != target {
					return false
				}
			}
			return true
		}

		for i := 0; i < len(buckets); i++ {
			if buckets[i]+nums[index] > target {
				fmt.Printf("---> continue index:%v buckets:%+v \n", index, buckets)
				continue
			}
			buckets[i] += nums[index]
			if backtrack(index + 1) {
				return true
			}
			buckets[i] -= nums[index]
		}
		return false
	}
	return backtrack(0)
}
