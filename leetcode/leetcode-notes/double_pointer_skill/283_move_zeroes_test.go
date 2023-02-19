package main

import (
	"reflect"
	"testing"
)

func TestMoveZeroes(t *testing.T) {
	type testCase struct {
		nums   []int
		expect []int
	}

	testCases := []testCase{
		{
			nums:   []int{0, 1, 0, 3, 12},
			expect: []int{1, 3, 12, 0, 0},
		},
		{
			nums:   []int{1},
			expect: []int{0},
		},
	}

	for _, tc := range testCases {
		moveZeroes(tc.nums)
		if reflect.DeepEqual(tc.nums, tc.expect) {
			t.Logf("result is %v,expect is %v", tc.nums, tc.expect)
		}
	}
}
