package main

import (
	"testing"
)

func TestLengthOfLIS(t *testing.T) {
	type testCase struct {
		nums   []int
		expect int
	}

	testCases := []testCase{
		{
			nums:   []int{10, 9, 2, 5, 3, 7, 101, 18},
			expect: 4,
		},
		{
			nums:   []int{0, 1, 0, 3, 2, 3},
			expect: 4,
		},
		{
			nums:   []int{7, 7, 7, 7, 7, 7, 7},
			expect: 1,
		},
	}

	for _, tc := range testCases {
		results := lengthOfLIS(tc.nums)
		if results != tc.expect {
			t.Fatalf("result is %v,expect is %v", results, tc.expect)
		}
	}
}
