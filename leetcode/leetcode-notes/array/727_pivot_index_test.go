package main

import "testing"

func TestPivotIndex(t *testing.T) {
	type testCase struct {
		nums   []int
		expect int
	}

	testCases := []testCase{
		{
			nums:   []int{1, 7, 3, 6, 5, 6},
			expect: 3,
		},
		{
			nums:   []int{1, 2, 3},
			expect: -1,
		},
		{
			nums:   []int{-1, -1, -1, -1, -1, 0},
			expect: 2,
		},
		{
			nums:   []int{2, 1, -1},
			expect: 0,
		},
	}

	for _, tc := range testCases {
		result := pivotIndex(tc.nums)
		if result != tc.expect {
			t.Logf("result is %v, expect is %v", result, tc.expect)
		}
	}
}
