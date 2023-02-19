package main

func pivotInteger(n int) int {
	left := 1
	right := n

	leftSum := left
	rightSum := right
	for left <= right {
		if leftSum == rightSum {
			if left == right {
				return left
			} else {
				left++
				leftSum += left
			}
		} else if leftSum < rightSum {
			left++
			leftSum += left
		} else if leftSum > rightSum {
			right--
			rightSum += right
		}
	}

	return -1
}
