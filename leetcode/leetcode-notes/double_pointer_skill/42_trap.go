package main

// O(n) + O(1)
func trap(height []int) int {
	var capSum int

	var lmax, rmax int
	left := 0
	right := len(height) - 1
	for left < right {
		lmax = max(lmax, height[left])
		rmax = max(rmax, height[right])

		if lmax < rmax {
			capSum += lmax - height[left]
			left++
		} else {
			capSum += rmax - height[right]
			right--
		}
	}
	return capSum
}

/*
// O(n) + O(n)
func trap(height []int) int {
	var capSum int

	lMax := make([]int, len(height))
	lMax[0] = height[0]
	for i := 1; i < len(height); i++ {
		lMax[i] = max(lMax[i-1], height[i])
	}

	rMax := make([]int, len(height))
	rMax[len(height)-1] = height[len(height)-1]
	for i := len(height) - 2; i >= 0; i-- {
		rMax[i] = max(rMax[i+1], height[i])
	}

	for i := 0; i < len(height); i++ {
		leftMax := lMax[i]
		rightMax := rMax[i]
		c := min(leftMax, rightMax) - height[i]
		capSum += c
	}
	return capSum

}
*/

func getMax(height []int) int {
	var max int
	for _, h := range height {
		if h > max {
			max = h
		}
	}
	return max
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}
