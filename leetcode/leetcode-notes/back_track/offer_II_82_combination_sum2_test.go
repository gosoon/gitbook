package main

import (
	"testing"
)

func TestCombinationSum2(t *testing.T) {
	type testCase struct {
		nums   []int
		target int
		expect [][]int
	}

	testCases := []testCase{
		{
			nums:   []int{10, 1, 2, 7, 6, 1, 5},
			target: 8,
			expect: [][]int{},
		},
		{
			nums:   []int{2, 5, 2, 1, 2},
			target: 5,
			expect: [][]int{},
		},
	}

	for _, tc := range testCases {
		results := combinationSum2(tc.nums, tc.target)
		t.Logf("result is %v,expect is %v", results, tc.expect)
	}
}
