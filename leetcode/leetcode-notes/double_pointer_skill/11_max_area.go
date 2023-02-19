package main

func maxArea(height []int) int {
	left := 0
	right := len(height) - 1
	var maxArea int
	for left < right {
		indexDiff := right - left
		var minHeight int
		if height[left] < height[right] {
			minHeight = height[left]
			left++
		} else {
			minHeight = height[right]
			right--
		}

		area := indexDiff * minHeight
		if area > maxArea {
			maxArea = area
		}
	}
	return maxArea
}
