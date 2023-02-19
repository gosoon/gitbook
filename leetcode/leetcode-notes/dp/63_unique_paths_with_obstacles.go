package main

import (
	"math"
)

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	row := len(obstacleGrid)
	col := len(obstacleGrid[0])
	if row == 1 && col == 1 {
		if obstacleGrid[0][0] == 0 {
			return 1
		}
		return 0
	}

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
			return 0
		}

		if obstacleGrid[i][j] == 1 {
			return 0
		}

		if i == 0 && j == 0 {
			return 1
		}

		if memo[i][j] != math.MaxInt32 {
			return memo[i][j]
		}

		memo[i][j] = dp(i, j-1) + dp(i-1, j)
		return memo[i][j]
	}

	return dp(row-1, col-1)
}
