package main

func nthUglyNumber(n int) int {

	checkUgly := func(n int) bool {
		if n == 1 {
			return true
		}
		for {
			if n%2 == 0 {
				n = n / 2
			} else if n%3 == 0 {
				n = n / 3
			} else if n%5 == 0 {
				n = n / 5
			} else if n == 1 {
				return true
			} else {
				return false
			}

		}
	}

	ct := 0
	for i := 1; ; {
		isUgly := checkUgly(i)
		if isUgly {
			ct++
		}
		if ct == n {
			return i
		}
		i++
	}

	return ct
}
