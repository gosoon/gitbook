package main

import "fmt"

func dividePlayers(skill []int) int64 {
	group := len(skill) / 2

	buckets,result := canPartitionKSubsets(skill, group)
	if !result  {
		return -1
	}

	for 

}

func canPartitionKSubsets(nums []int, k int) ([]int, bool) {
	if k > len(nums) {
		return []int(nil), false
	}

	var sum int
	for _, num := range nums {
		sum += num
	}
	if sum%k != 0 {
		return []int(nil), false
	}

	var backtrack func(index int) bool

	target := sum / k
	buckets := make([]int, k)
	bucketsList := make([][]int,k)
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
	result := backtrack(0)
	return buckets, result
}
