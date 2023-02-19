package main

func minAddToMakeValid(s string) int {
	var need, left int
	for _, b := range s {
		if b == '(' {
			left++
		} else {
			if left == 0 {
				need++
			} else {
				left--
			}
		}
	}
	return need + left
}
