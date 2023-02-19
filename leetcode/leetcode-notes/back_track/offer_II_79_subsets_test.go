package main

import (
	"testing"
)

func TestSubsets(t *testing.T) {
	type testCase struct {
		nums   []int
		expect [][]int
	}

	testCases := []testCase{
		{
			nums:   []int{1, 2, 3},
			expect: [][]int{},
		},
	}

	for _, tc := range testCases {
		results := subsets(tc.nums)
		t.Logf("result is %v,expect is %v", results, tc.expect)
	}
}
