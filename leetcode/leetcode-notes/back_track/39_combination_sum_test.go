package main

import (
	"testing"
)

func TestCombinationSum(t *testing.T) {
	type testCase struct {
		nums   []int
		target int
		expect [][]int
	}

	testCases := []testCase{
		{
			nums:   []int{2, 3, 6, 7},
			target: 7,
			expect: [][]int{},
		},
		{
			nums:   []int{2, 3, 5},
			target: 8,
			expect: [][]int{},
		},
		{
			nums:   []int{2},
			target: 1,
			expect: [][]int{},
		},
	}

	for _, tc := range testCases {
		results := combinationSum(tc.nums, tc.target)
		t.Logf("result is %v,expect is %v", results, tc.expect)
	}
}
