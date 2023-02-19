package main

func maximumCount(nums []int) int {
	var pos, neg int
	for _, num := range nums {
		if num > 0 {
			pos++
		} else if num < 0 {
			neg++
		}
	}
	if pos > neg {
		return pos
	}
	return neg
}
