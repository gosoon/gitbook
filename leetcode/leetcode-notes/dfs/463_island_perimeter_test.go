package dfs

import (
	"testing"
)

func TestIslandPerimeter(t *testing.T) {
	type testCase struct {
		grid   [][]int
		expect int
	}

	testCases := []testCase{
		{
			grid: [][]int{
				{0, 1, 0, 0}, {1, 1, 1, 0}, {0, 1, 0, 0}, {1, 1, 0, 0},
			},
			expect: 16,
		},
		//{
		//grid: [][]int{
		//{1},
		//},
		//expect: 4,
		//},
	}

	for _, tc := range testCases {
		results := islandPerimeter(tc.grid)
		if results != tc.expect {
			t.Logf("result is %v,expect is %v", results, tc.expect)
		}
	}
}
