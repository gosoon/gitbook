package main

import (
	"testing"
)

func TestFindSubsequences(t *testing.T) {
	type testCase struct {
		nums   []int
		expect [][]int
	}

	testCases := []testCase{
		{
			nums:   []int{4, 6, 7, 7},
			expect: [][]int{},
		},
		{
			nums:   []int{4, 4, 3, 2, 1},
			expect: [][]int{},
		},
	}

	for _, tc := range testCases {
		results := findSubsequences(tc.nums)
		t.Logf("result is %v", results)
	}
}
