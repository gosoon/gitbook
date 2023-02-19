package main

import "testing"

func TestTrap(t *testing.T) {
	type testCase struct {
		nums   []int
		expect int
	}

	testCases := []testCase{
		{
			nums:   []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1},
			expect: 6,
		},
		{
			nums:   []int{4, 2, 0, 3, 2, 5},
			expect: 9,
		},
	}

	for _, tc := range testCases {
		t.Logf("nums is %v", tc.nums)
		result := trap(tc.nums)
		if result != tc.expect {
			t.Logf("result is %v,expect is %v", result, tc.expect)
		}
	}
}
