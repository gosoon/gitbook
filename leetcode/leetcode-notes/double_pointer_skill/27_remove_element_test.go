package main

import "testing"

func TestRemoveElement(t *testing.T) {
	type testCase struct {
		nums   []int
		val    int
		expect int
	}

	testCases := []testCase{
		{
			nums:   []int{3, 2, 2, 3},
			val:    3,
			expect: 2,
		},
		{
			nums:   []int{0, 1, 2, 2, 3, 0, 4, 2},
			val:    2,
			expect: 5,
		},
	}

	for _, tc := range testCases {
		result := removeElement(tc.nums, tc.val)
		if result != tc.expect {
			t.Logf("result is %v,expect is %v", result, tc.expect)
		}
	}
}
