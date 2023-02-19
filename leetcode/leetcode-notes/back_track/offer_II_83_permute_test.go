package main

import (
	"testing"
)

func TestPermute(t *testing.T) {
	type testCase struct {
		nums   []int
		expect [][]int
	}

	testCases := []testCase{
		{
			nums: []int{1, 2, 3},
		},
	}

	for _, tc := range testCases {
		results := permute(tc.nums)
		t.Logf("result is %v,expect is %v", results, tc.expect)
	}
}
