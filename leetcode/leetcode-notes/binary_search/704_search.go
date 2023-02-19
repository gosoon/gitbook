package main

import "fmt"

func search(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	for left <= right {
		mid := left + (right-left)/2
		fmt.Printf("left:%v,right:%v,mid:%v\n", left, right, mid)
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			right = mid - 1
		} else if nums[mid] < target {
			left = mid + 1
		}
		fmt.Printf("-->left:%v,right:%v,mid:%v\n", left, right, mid)
	}

	return -1
}
