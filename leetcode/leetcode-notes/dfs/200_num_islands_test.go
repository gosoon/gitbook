package dfs

import (
	"testing"
)

func TestNumIslands(t *testing.T) {
	type testCase struct {
		grid   [][]byte
		expect int
	}

	testCases := []testCase{
		{
			grid: [][]byte{
				{'1', '1', '0', '0', '0'},
				{'1', '1', '0', '0', '0'},
				{'0', '0', '1', '0', '0'},
				{'0', '0', '0', '1', '1'},
			},
			expect: 3,
		},
	}

	for _, tc := range testCases {
		results := numIslands(tc.grid)
		if results != tc.expect {
			t.Logf("result is %v,expect is %v", results, tc.expect)
		}
	}
}
