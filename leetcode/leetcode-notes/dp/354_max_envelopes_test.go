package main

import (
	"testing"
)

func TestMaxEnvelopes(t *testing.T) {
	type testCase struct {
		nums   [][]int
		expect int
	}

	testCases := []testCase{
		{
			nums:   [][]int{{5, 4}, {6, 4}, {6, 7}, {2, 3}},
			expect: 3,
		},
		{
			nums:   [][]int{{1, 1}, {1, 1}, {1, 1}},
			expect: 1,
		},
	}

	for _, tc := range testCases {
		results := maxEnvelopes(tc.nums)
		if results != tc.expect {
			t.Fatalf("result is %v,expect is %v", results, tc.expect)
		}
	}
}
