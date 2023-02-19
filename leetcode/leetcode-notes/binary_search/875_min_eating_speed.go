package main

func minEatingSpeed(piles []int, h int) int {
	left := 1
	right := 1000000000

	for left <= right {
		mid := left + (right-left)/2
		hours := f(piles, mid)
		if hours <= h {
			right = mid - 1
		} else if hours > h {
			left = mid + 1
		}
	}

	return left
}

func f(piles []int, x int) int {
	if x == 0 {
		return 0
	}

	var hours int
	for i := 0; i < len(piles); i++ {
		hour := piles[i] / x
		hours += hour
		if piles[i]%x != 0 {
			hours += 1
		}
	}
	return hours
}
