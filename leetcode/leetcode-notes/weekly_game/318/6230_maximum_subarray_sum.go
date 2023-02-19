package main

import "fmt"

func maximumSubarraySum(nums []int, k int) int64 {
	left, right := 0, 0
	var sum int
	var maximum int
	m := make(map[int]int)
	for right < len(nums) {
		fmt.Printf("left:%v right: %v \n", left, right)
		// scale up
		num := nums[right]
		right++

		m[num]++
		sum += num
		count := m[num]
		// handle repeat
		if count > 1 {
			fmt.Printf("left:%v right: %v found:%v \n", left, right, num)
			for left < right {
				n := nums[left]
				left++
				sum -= n
				m[n]--
				if n == num {
					break
				}
			}
		}
		// scale down
		for right-left > k {
			fmt.Printf(">k, left:%v right: %v \n", left, right)
			num := nums[left]
			left++
			sum -= num
			m[num]--
		}

		if right-left == k {
			fmt.Printf("== k, left:%v right: %v \n", left, right)
			maximum = max(maximum, sum)
		}
	}
	return int64(maximum)
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
