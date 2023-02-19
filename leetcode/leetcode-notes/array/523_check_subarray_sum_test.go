package main

import "testing"

func TestCheckSubarraySum(t *testing.T) {
	type testCase struct {
		nums []int
		k    int
	}

	testCases := []testCase{
		{
			nums: []int{23, 2, 4, 6, 6},
			k:    7,
		},
		{
			nums: []int{23, 2, 6, 4, 7},
			k:    6,
		},
		{
			nums: []int{23, 2, 6, 4, 7},
			k:    13,
		},
		{
			nums: []int{1, 2, 12},
			k:    6,
		},
		{
			nums: []int{5, 0, 0, 0},
			k:    3,
		},
		{
			nums: []int{-1, 2},
			k:    2,
		},
	}

	for _, tc := range testCases {
		t.Log(checkSubarraySum(tc.nums, tc.k))
	}
}
