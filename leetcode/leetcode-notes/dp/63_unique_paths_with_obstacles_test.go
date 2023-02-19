package main

import (
	"testing"
)

func TestUniquePathsWithObstacles(t *testing.T) {
	type testCase struct {
		obstacleGrid [][]int
		expect       int
	}

	testCases := []testCase{
		{
			obstacleGrid: [][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}},
			expect:       2,
		},
		{
			obstacleGrid: [][]int{{0, 1}, {0, 0}},
			expect:       1,
		},
		{
			obstacleGrid: [][]int{{0}},
			expect:       1,
		},
		{
			obstacleGrid: [][]int{{1}},
			expect:       0,
		},
		{
			obstacleGrid: [][]int{{0, 1}},
			expect:       0,
		},
		{
			obstacleGrid: [][]int{{1, 0}},
			expect:       0,
		},
	}

	for _, tc := range testCases {
		results := uniquePathsWithObstacles(tc.obstacleGrid)
		if results != tc.expect {
			t.Fatalf("result is %v,expect is %v", results, tc.expect)
		}
	}
}
