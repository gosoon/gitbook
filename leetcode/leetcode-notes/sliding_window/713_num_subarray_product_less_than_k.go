package main

import "fmt"

func numSubarrayProductLessThanK(nums []int, k int) int {

	var l, r int
	prod := 1
	var result int
	for r < len(nums) {
		// scale up
		n := nums[r]
		r++
		prod *= n

		if prod < k {
			result++
		}
		fmt.Printf("l:%v,r:%v,result:%v,prod:%v \n", l, r, result, prod)
		for l < r && prod >= k {
			// scale down
			n := nums[l]
			l++
			prod /= n

			if prod < k {
				result++
			}
			fmt.Printf("-->l:%v,r:%v,result:%v,prod:%v \n", l, r, result, prod)
		}
	}
	return result
}
