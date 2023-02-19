package main

import (
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	type testCase struct {
		nums   []int
		target int
		expect []int
	}

	testCases := []testCase{
		{
			nums:   []int{2, 7, 11, 15},
			target: 9,
			expect: []int{1, 2},
		},
		{
			nums:   []int{2, 3, 4},
			target: 6,
			expect: []int{1, 3},
		},
		{
			nums:   []int{-1, 0},
			target: -1,
			expect: []int{1, 2},
		},
	}

	for _, tc := range testCases {
		result := twoSum(tc.nums, tc.target)
		if !reflect.DeepEqual(result, tc.expect) {
			t.Logf("result is %v,expect is %v", result, tc.expect)
		}
	}
}
