package main

import "fmt"

func peakIndexInMountainArray(arr []int) int {
	left := 1
	right := len(arr) - 1
	for left < right {
		// mid := left + (right-left)/2
		mid := (left + right + 1) >> 1
		fmt.Printf("left:%v,right:%v,mid:%v\n", left, right, mid)
		if arr[mid-1] < arr[mid] {
			left = mid
		} else {
			right = mid - 1
		}
	}
	return left
}

/*
func peakIndexInMountainArray(arr []int) int {
	left := 0
	right := len(arr) - 1

	for left <= right {
		mid := left + (right-left)/2
		fmt.Printf("left:%v,right:%v,mid:%v\n", left, right, mid)
		if mid == 0 {
			if arr[mid] > arr[mid+1] {
				return mid
			} else {
				left = mid + 1
			}
		} else if mid == len(arr)-1 {
			if arr[mid] > arr[mid-1] {
				return mid
			} else {
				right = mid - 1
			}
		} else {
			if arr[mid] > arr[mid-1] && arr[mid] > arr[mid+1] {
				return mid
			} else if arr[mid] > arr[mid-1] && arr[mid] < arr[mid+1] {
				left = mid + 1
			} else if arr[mid] < arr[mid-1] && arr[mid] > arr[mid+1] {
				right = mid - 1
			}
		}
	}

	return -1
}
*/
