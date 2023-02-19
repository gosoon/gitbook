package main

import (
	"math"
)

func minFallingPathSum(matrix [][]int) int {
	row := len(matrix)
	if row == 0 {
		return matrix[0][0]
	}
	col := len(matrix[0])

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
			return matrix[i][j]
		}

		if memo[i][j] != math.MaxInt32 {
			return memo[i][j]
		}

		memo[i][j] = matrix[i][j] + minFuncByMinFallingPathSum(dp(i-1, j),
			minFuncByMinFallingPathSum(dp(i-1, j-1), dp(i-1, j+1)))
		return memo[i][j]
	}

	result := math.MaxInt32
	for j := 0; j < col; j++ {
		result = minFuncByMinFallingPathSum(result, dp(row-1, j))
	}

	return result
}
func minFuncByMinFallingPathSum(x, y int) int {
	if x < y {
		return x
	}
	return y
}
