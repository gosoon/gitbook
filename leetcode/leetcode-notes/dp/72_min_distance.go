package main

import "math"

func minDistance(word1 string, word2 string) int {
	m := len(word1)
	n := len(word2)
	memo := make([][]int, m)
	for i := range memo {
		memo[i] = make([]int, n)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			memo[i][j] = math.MaxInt64
		}
	}
	var dp func(s1 string, i int, s2 string, j int) int
	dp = func(s1 string, i int, s2 string, j int) int {
		if i == -1 {
			return j + 1
		}
		if j == -1 {
			return i + 1
		}

		if memo[i][j] != math.MaxInt64 {
			return memo[i][j]
		}

		if s1[i] == s2[j] {
			memo[i][j] = dp(s1, i-1, s2, j-1)
			return memo[i][j]
		}
		// insert,delete,replace
		memo[i][j] = minForMinDistance(dp(s1, i, s2, j-1)+1,
			dp(s1, i-1, s2, j)+1,
			dp(s1, i-1, s2, j-1)+1,
		)
		return memo[i][j]
	}

	return dp(word1, m-1, word2, n-1)
}

func minForMinDistance(x, y, z int) int {
	if x < y {
		if x < z {
			return x
		}
		return z
	} else {
		if y < z {
			return y
		}
		return z
	}
}
