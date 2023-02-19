package main

func removeElement(nums []int, val int) int {
	i, j := 0, 0
	for j < len(nums) {
		if nums[j] == val {
			j++
		} else {
			nums[i] = nums[j]
			i++
			j++
		}
	}
	return i
}
