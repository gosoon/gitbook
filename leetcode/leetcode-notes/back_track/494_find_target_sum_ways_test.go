package main

import (
	"testing"
)

func TestFindTargetSumWays(t *testing.T) {
	type testCase struct {
		nums   []int
		target int
		expect int
	}

	testCases := []testCase{
		//{
		//nums:   []int{1, 1, 1, 1, 1},
		//target: 3,
		//expect: 5,
		//},
		//{
		//nums:   []int{1},
		//target: 1,
		//expect: 1,
		//},
		//{
		//nums:   []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
		//target: 1,
		//expect: 1,
		//},
		//{
		//nums:   []int{0, 0, 0, 0, 0, 0, 0, 0, 1},
		//target: 1,
		//expect: 256,
		//},
		{
			nums:   []int{42, 24, 30, 14, 38, 27, 12, 29, 43, 42, 5, 18, 0, 1, 12, 44, 45, 50, 21, 47},
			target: 38,
			expect: 5602,
		},
	}

	for _, tc := range testCases {
		results := findTargetSumWays(tc.nums, tc.target)
		if results != tc.expect {
			t.Logf("result is %v,expect is %v", results, tc.expect)
		}
	}
}
