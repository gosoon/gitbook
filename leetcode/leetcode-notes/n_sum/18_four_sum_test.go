package main

import "testing"

func TestFourSum(t *testing.T) {
	type testCase struct {
		nums   []int
		target int
	}

	testCases := []testCase{
		{
			nums:   []int{1, 0, -1, 0, -2, 2},
			target: 0,
		},
		{
			nums:   []int{2, 2, 2, 2, 2},
			target: 8,
		},
		{
			nums:   []int{},
			target: 0,
		},
	}

	for _, tc := range testCases {
		t.Log(fourSum(tc.nums, tc.target))
		//t.Log(twoSumCommonFunc(tc.nums, 0, tc.target))
	}
}
