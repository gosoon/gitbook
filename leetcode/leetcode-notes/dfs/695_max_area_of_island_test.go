package dfs

import (
	"testing"
)

func TestMaxAreaOfIsland(t *testing.T) {
	type testCase struct {
		grid   [][]int
		expect int
	}

	testCases := []testCase{
		{
			grid: [][]int{
				{0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
				{0, 1, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 1, 0, 0, 1, 1, 0, 0, 1, 0, 1, 0, 0},
				{0, 1, 0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0},
			},
			expect: 6,
		},
		{
			grid: [][]int{
				{0, 0, 0, 0, 0, 0, 0, 0},
			},
			expect: 0,
		},
	}

	for _, tc := range testCases {
		results := maxAreaOfIsland(tc.grid)
		if results != tc.expect {
			t.Logf("result is %v,expect is %v", results, tc.expect)
		}
	}
}
