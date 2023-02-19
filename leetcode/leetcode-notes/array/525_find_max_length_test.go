package main

import "testing"

func TestFindMaxLength(t *testing.T) {
	type testCase struct {
		nums []int
	}

	testCases := []testCase{
		{
			nums: []int{0, 1},
		},
		{
			nums: []int{0, 1, 0},
		},
	}

	for _, tc := range testCases {
		t.Log(findMaxLength(tc.nums))
	}
}
