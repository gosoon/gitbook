package main

import "fmt"

func moveZeroes(nums []int) {
	i, j := 0, 0

	for j < len(nums) {
		if nums[j] == 0 {
			j++
		} else {
			nums[i] = nums[j]
			if i != j {
				nums[j] = 0
			}
			i++
			j++
		}
	}
	fmt.Println(nums)
}
