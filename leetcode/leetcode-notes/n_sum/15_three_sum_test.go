package main

import "testing"

func TestThreeSum(t *testing.T) {
	type testCase struct {
		nums []int
	}

	testCases := []testCase{
		{
			nums: []int{-1, 0, 1, 2, -1, -4},
		},
		{
			nums: []int{0, 1, 1},
		},
		{
			nums: []int{},
		},
		{
			nums: []int{-2, 0, 1, 1, 2},
		},
	}

	for _, tc := range testCases {
		t.Log(threeSum(tc.nums))
	}
}
