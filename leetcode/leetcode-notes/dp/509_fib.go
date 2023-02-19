package main

func fib(n int) int {
	if n < 2 {
		return n
	}
	dp0 := 0
	dp1 := 1
	var dp2 int
	for i := 2; i <= n; i++ {
		dp2 = dp0 + dp1
		dp0 = dp1
		dp1 = dp2
	}
	return dp2
}
