package dfs

import (
	"testing"
)

func TestCountSubIslands(t *testing.T) {
	type testCase struct {
		grid1  [][]int
		grid2  [][]int
		expect int
	}

	testCases := []testCase{
		{
			grid1: [][]int{
				{1, 1, 1, 0, 0},
				{0, 1, 1, 1, 1},
				{0, 0, 0, 0, 0},
				{1, 0, 0, 0, 0},
				{1, 1, 0, 1, 1},
			},
			grid2: [][]int{
				{1, 1, 1, 0, 0}, {0, 0, 1, 1, 1}, {0, 1, 0, 0, 0}, {1, 0, 1, 1, 0}, {0, 1, 0, 1, 0},
			},
			expect: 3,
		},
	}

	for _, tc := range testCases {
		results := countSubIslands(tc.grid1, tc.grid2)
		if results != tc.expect {
			t.Logf("result is %v,expect is %v", results, tc.expect)
		}
	}
}
