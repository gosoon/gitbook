package main

import (
	"testing"
)

func TestCombinationSum4(t *testing.T) {
	type testCase struct {
		nums   []int
		target int
		expect [][]int
	}

	testCases := []testCase{
		{
			nums:   []int{1, 2, 3},
			target: 32,
			expect: [][]int{},
		},
	}

	for _, tc := range testCases {
		results := combinationSum4(tc.nums, tc.target)
		t.Logf("result is %v,expect is %v", results, tc.expect)
	}
}
