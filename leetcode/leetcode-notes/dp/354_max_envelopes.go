package main

import (
	"fmt"
	"sort"
)

func maxEnvelopes(envelopes [][]int) int {
	sort.Slice(envelopes, func(i, j int) bool {
		if envelopes[i][0] < envelopes[j][0] {
			return true
		} else if envelopes[i][0] == envelopes[j][0] {
			if envelopes[i][1] > envelopes[j][1] {
				return true
			}
		}
		return false
	})

	fmt.Printf("%+v \n", envelopes)

	var nums []int
	for i := 0; i < len(envelopes); i++ {
		nums = append(nums, envelopes[i][1])
	}

	fmt.Println(nums)
	res := lengthOfLIS(nums)
	return res
}
