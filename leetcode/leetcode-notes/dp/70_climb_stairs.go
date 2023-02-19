package main

// f(x) = f(x-1) + f(x-2)
func climbStairs(n int) int {
	if n <= 2 {
		return n
	}

	var dp1, dp2, dpn int
	dp1 = 1
	dp2 = 2
	for i := 3; i <= n; i++ {
		dpn = dp1 + dp2
		dp1 = dp2
		dp2 = dpn
	}
	return dpn
}
