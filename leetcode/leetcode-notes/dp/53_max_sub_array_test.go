package main

import (
	"testing"
)

func TestMaxSubArray(t *testing.T) {
	type testCase struct {
		nums   []int
		expect int
	}

	testCases := []testCase{
		{
			nums:   []int{-2, 1, -3, 4, -1, 2, 1, -5, 4},
			expect: 6,
		},
		{
			nums:   []int{1},
			expect: 1,
		},
		{
			nums:   []int{5, 4, -1, 7, 8},
			expect: 23,
		},
	}

	for _, tc := range testCases {
		results := maxSubArray(tc.nums)
		if results != tc.expect {
			t.Fatalf("result is %v,expect is %v", results, tc.expect)
		}
	}
}
