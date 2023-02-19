package main

import "testing"

func TestRemoveDuplicates(t *testing.T) {
	type testCase struct {
		nums   []int
		expect int
	}

	testCases := []testCase{
		{
			nums:   []int{1, 1, 2},
			expect: 2,
		},
		{
			nums:   []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4},
			expect: 5,
		},
	}

	for _, tc := range testCases {
		result := removeDuplicates(tc.nums)
		if result != tc.expect {
			t.Logf("result is %v,expect is %v", result, tc.expect)
		}
	}
}
