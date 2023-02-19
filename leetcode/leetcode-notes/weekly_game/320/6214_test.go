package main

import (
	"testing"
)

func TestUnequalTriplets(t *testing.T) {
	type testCase struct {
		nums   []int
		expect int
	}

	testCases := []testCase{
		{
			nums:   []int{4, 4, 2, 4, 3},
			expect: 3,
		},
		{
			nums:   []int{1, 1, 1, 1, 1},
			expect: 0,
		},
		{
			nums:   []int{1, 3, 1, 2, 4},
			expect: 7,
		},
	}

	for _, tc := range testCases {
		results := unequalTriplets(tc.nums)
		t.Logf("result is %v,expect is %v", results, tc.expect)
	}
}
