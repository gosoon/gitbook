package main

import (
	"fmt"
	"sort"
)

func searchByOffer(nums []int, target int) int {
	leftIndex := sort.SearchInts(nums, target)
	fmt.Println(leftIndex)
	if len(nums) == 0 || leftIndex >= len(nums) || nums[leftIndex] != target {
		return 0
	}
	rightIndex := sort.SearchInts(nums, target+1)
	fmt.Println(rightIndex)
	return rightIndex - leftIndex
}
