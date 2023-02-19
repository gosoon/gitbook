package main

import (
	"testing"
)

func TestMinFallingPathSum(t *testing.T) {
	type testCase struct {
		matrix [][]int
		expect int
	}

	testCases := []testCase{
		{
			matrix: [][]int{{2, 1, 3}, {6, 5, 4}, {7, 8, 9}},
			expect: 13,
		},
		{
			matrix: [][]int{{-19, 57}, {-40, -5}},
			expect: -59,
		},
	}

	for _, tc := range testCases {
		results := minFallingPathSum(tc.matrix)
		if results != tc.expect {
			t.Fatalf("result is %v,expect is %v", results, tc.expect)
		}
	}
}
