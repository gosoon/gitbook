package main

import "testing"

func TestMaxArea(t *testing.T) {
	type testCase struct {
		nums   []int
		expect int
	}

	testCases := []testCase{
		{
			nums:   []int{1, 8, 6, 2, 5, 4, 8, 3, 7},
			expect: 49,
		},
		{
			nums:   []int{1, 1},
			expect: 1,
		},
	}

	for _, tc := range testCases {
		t.Logf("nums is %v", tc.nums)
		result := maxArea(tc.nums)
		if result != tc.expect {
			t.Logf("result is %v,expect is %v", result, tc.expect)
		}
	}
}
