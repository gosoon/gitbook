package main

import (
	"testing"
)

func TestMinCost(t *testing.T) {
	type testCase struct {
		costs  [][]int
		expect int
	}

	testCases := []testCase{
		{
			costs:  [][]int{{17, 2, 17}, {16, 16, 5}, {14, 3, 19}},
			expect: 10,
		},
		{
			costs:  [][]int{{7, 6, 2}},
			expect: 2,
		},
	}

	for _, tc := range testCases {
		results := minCost(tc.costs)
		if results != tc.expect {
			t.Fatalf("result is %v,expect is %v", results, tc.expect)
		}
	}
}
