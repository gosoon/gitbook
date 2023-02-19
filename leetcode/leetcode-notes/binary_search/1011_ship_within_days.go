package main

import "fmt"

func shipWithinDays(weights []int, days int) int {
	left := 0
	right := 0

	for _, weight := range weights {
		if weight > left {
			left = weight
		}
		right += weight
	}

	fmt.Println(left, right)
	for left <= right {
		mid := left + (right-left)/2
		needDays := getNeedDays(weights, mid)
		fmt.Printf("left:%v right:%v mid:%v needDays:%v \n", left, right, mid, needDays)
		if needDays == days {
			right = mid - 1
		} else if needDays < days {
			right = mid - 1
		} else if needDays > days {
			left = mid + 1
		}
	}

	return left
}
func getNeedDays(weights []int, k int) int {
	var days int
	for i := 0; i < len(weights); {
		caps := k
		for i < len(weights) {
			if caps < weights[i] {
				break
			}
			caps -= weights[i]
			i++
		}
		days++
	}

	return days
}

/*
// f(x)
func getNeedDays(weights []int, k int) int {
	var caps int
	var days int
	for i := 0; i < len(weights); i++ {
		caps += weights[i]
		if caps > k {
			days++
			caps = 0
			if weights[i] >= k {
				days += weights[i] / k
				if weights[i]%k != 0 {
					days++
				}
			} else {
				caps += weights[i]
			}
		}
		//fmt.Printf("i:%v caps:%v days:%v \n", i, caps, days)
	}
	days += caps / k
	if caps%k != 0 {
		days++
	}

	return days
}
*/
