package main

import "testing"

func TestSubarraysDivByK(t *testing.T) {
	type testCase struct {
		nums   []int
		k      int
		expect int
	}

	testCases := []testCase{
		{
			nums:   []int{4, 5, 0, -2, -3, 1},
			k:      5,
			expect: 7,
		},
		{
			nums:   []int{5},
			k:      9,
			expect: 0,
		},
		{
			nums:   []int{-1, 2, 9},
			k:      2,
			expect: 2,
		},
	}

	for _, tc := range testCases {
		result := subarraysDivByK(tc.nums, tc.k)
		if tc.expect != result {
			t.Logf("not pass,result is %v,expect is %v", result, tc.expect)
		}
	}
}
