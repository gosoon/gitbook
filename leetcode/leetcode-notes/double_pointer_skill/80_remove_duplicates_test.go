package main

import "testing"

func TestRemoveDuplicatesII(t *testing.T) {
	type testCase struct {
		nums   []int
		expect int
	}

	testCases := []testCase{
		{
			nums:   []int{1, 1, 1, 2, 2, 3},
			expect: 5,
		},
		{
			nums:   []int{0, 0, 1, 1, 1, 1, 2, 3, 3},
			expect: 7,
		},
	}

	for _, tc := range testCases {
		t.Logf("nums is %v", tc.nums)
		result := removeDuplicatesII(tc.nums)
		if result != tc.expect {
			t.Logf("result is %v,expect is %v", result, tc.expect)
		}
	}
}
