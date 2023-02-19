package main

import (
	"fmt"
	"sort"
)

func findMaxK(nums []int) int {
	left := 0
	right := len(nums) - 1

	sort.Ints(nums)
	fmt.Println(nums)
	for left < right {
		fmt.Printf("left:%v right:%v \n", nums[left], nums[right])

		if nums[left] == -nums[right] {
			return nums[right]
		} else if nums[left] > -nums[right] {
			right--
		} else {
			left++
		}
	}

	return -1
}
