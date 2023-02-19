package main

import "fmt"

func removeDuplicates(nums []int) int {
	var slow, fast int
	for fast < len(nums) {
		if nums[fast] == nums[slow] {
			fast++
		} else {
			slow++
			nums[slow] = nums[fast]
			fast++
		}
		fmt.Println(nums)
	}
	fmt.Println(nums)
	return slow + 1
}
