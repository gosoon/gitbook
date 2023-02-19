package main

import "fmt"

func removeDuplicatesII(nums []int) int {
	var slow, fast int

	for ; fast < len(nums); fast++ {
		if fast < 2 || nums[fast] != nums[slow-2] {
			nums[slow] = nums[fast]
			slow++

		}
	}

	fmt.Printf("slow is %v,fast is %v,nums is %v \n", slow, fast, nums)
	return slow
}
