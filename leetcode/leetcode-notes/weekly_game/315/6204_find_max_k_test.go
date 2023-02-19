package main

import (
	"testing"
)

func TestFindMaxK(t *testing.T) {
	type testCase struct {
		nums   []int
		expect int
	}

	testCases := []testCase{
		{
			nums:   []int{-1, 2, -3, 3},
			expect: 3,
		},
		{
			nums:   []int{-1, 10, 6, 7, -7, 1},
			expect: 7,
		},
		{
			nums:   []int{-10, 8, 6, 7, -2, -3},
			expect: -1,
		},
	}

	for _, tc := range testCases {
		results := findMaxK(tc.nums)
		if results != tc.expect {
			t.Logf("result is %v,expect is %v", results, tc.expect)
		}
	}
}
