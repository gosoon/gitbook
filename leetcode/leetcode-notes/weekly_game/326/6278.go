package main

func countDigits(num int) int {
	var result int
	t := num
	for t != 0 {
		if t%10 != 0 {
			k := t % 10
			t = t / 10
			if num%k == 0 {
				result++
			}
		}
	}
	return result
}
