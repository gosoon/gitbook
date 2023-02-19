package main

import "testing"

func TestTwoSumUpgrade(t *testing.T) {
	type testCase struct {
		nums   []int
		target int
	}

	testCases := []testCase{
		{
			nums:   []int{2, 7, 11, 15},
			target: 9,
		},
		{
			nums:   []int{3, 2, 4},
			target: 6,
		},
		{
			nums:   []int{3, 3},
			target: 6,
		},
		{
			nums:   []int{},
			target: 0,
		},
	}

	for _, tc := range testCases {
		t.Log(twoSumUpgrade(tc.nums, tc.target))
	}
}
