package dfs

import (
	"testing"
)

func TestClosedIsland(t *testing.T) {
	type testCase struct {
		grid   [][]int
		expect int
	}

	testCases := []testCase{
		{
			grid: [][]int{
				{1, 1, 1, 1, 1, 1, 1, 0},
				{1, 0, 0, 0, 0, 1, 1, 0},
				{1, 0, 1, 0, 1, 1, 1, 0},
				{1, 0, 0, 0, 0, 1, 0, 1},
				{1, 1, 1, 1, 1, 1, 1, 0},
			},
			expect: 2,
		},
		{
			grid: [][]int{
				{0, 0, 1, 1, 0, 1, 0, 0, 1, 0},
				{1, 1, 0, 1, 1, 0, 1, 1, 1, 0},
				{1, 0, 1, 1, 1, 0, 0, 1, 1, 0},
				{0, 1, 1, 0, 0, 0, 0, 1, 0, 1},
				{0, 0, 0, 0, 0, 0, 1, 1, 1, 0},
				{0, 1, 0, 1, 0, 1, 0, 1, 1, 1},
				{1, 0, 1, 0, 1, 1, 0, 0, 0, 1},
				{1, 1, 1, 1, 1, 1, 0, 0, 0, 0},
				{1, 1, 1, 0, 0, 1, 0, 1, 0, 1},
				{1, 1, 1, 0, 1, 1, 0, 1, 1, 0},
			},
			expect: 5,
		},
	}

	for _, tc := range testCases {
		results := closedIsland(tc.grid)
		if results != tc.expect {
			t.Logf("result is %v,expect is %v", results, tc.expect)
		}
	}
}
