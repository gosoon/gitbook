package main

import "fmt"

func searchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}
	left := findLeftPosition(nums, target)
	if left == -1 {
		return []int{-1, -1}
	}
	right := findRightPosition(nums, target)
	return []int{left, right}
}

func findLeftPosition(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	for left <= right {
		mid := left + (right-left)/2
		fmt.Printf("left: left:%v,right:%v,mid:%v \n", left, right, mid)
		if nums[mid] == target {
			right = mid - 1
		} else if nums[mid] > target {
			right = mid - 1
		} else if nums[mid] < target {
			left = mid + 1
		}
	}
	if left == len(nums) {
		return -1
	}

	if nums[left] == target {
		return left
	}

	return -1
}

func findRightPosition(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	for left <= right {
		mid := left + (right-left+1)/2
		fmt.Printf("right: left:%v,right:%v,mid:%v \n", left, right, mid)
		if nums[mid] == target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid - 1
		} else if nums[mid] < target {
			left = mid + 1
		}
	}

	if nums[left-1] == target {
		return left - 1
	}

	return -1
}
