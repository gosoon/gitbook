package main

import (
	"math"
)

func minCost(costs [][]int) int {
	row := len(costs)
	if row == 0 {
		return -1
	}
	col := len(costs[0])

	memo := make([][]int, row)
	for i := 0; i < row; i++ {
		memo[i] = make([]int, col)
		for j := 0; j < col; j++ {
			memo[i][j] = math.MaxInt32
		}
	}

	var dp func(i int, j int) int
	dp = func(i int, j int) int {
		if i < 0 || i >= row || j < 0 || j >= col {
			return math.MaxInt32
		}

		if i == 0 {
			return costs[i][j]
		}

		if memo[i][j] != math.MaxInt32 {
			return memo[i][j]
		}

		if j == 0 {
			memo[i][j] = costs[i][j] + minFuncByMinCost(dp(i-1, j+1), dp(i-1, j+2))
		} else if j == 1 {
			memo[i][j] = costs[i][j] + minFuncByMinCost(dp(i-1, j-1), dp(i-1, j+1))
		} else if j == 2 {
			memo[i][j] = costs[i][j] + minFuncByMinCost(dp(i-1, j-1), dp(i-1, j-2))
		}
		return memo[i][j]
	}

	result := math.MaxInt32
	for j := 0; j < col; j++ {
		result = minFuncByMinCost(result, dp(row-1, j))
	}

	return result
}
func minFuncByMinCost(x, y int) int {
	if x < y {
		return x
	}
	return y
}
