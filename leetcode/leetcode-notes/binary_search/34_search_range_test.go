package main

import (
	"reflect"
	"testing"
)

func TestSearchRange(t *testing.T) {
	type testCase struct {
		nums   []int
		target int
		expect []int
	}

	testCases := []testCase{
		{
			nums:   []int{5, 7, 7, 8, 8, 10},
			target: 8,
			expect: []int{3, 4},
		},
		{
			nums:   []int{5, 7, 7, 8, 8, 10},
			target: 6,
			expect: []int{-1, -1},
		},
		{
			nums:   []int{2, 2},
			target: 2,
			expect: []int{0, 1},
		},
		{
			nums:   []int{2, 2},
			target: 3,
			expect: []int{-1, -1},
		},
	}

	for _, tc := range testCases {
		result := searchRange(tc.nums, tc.target)
		if !reflect.DeepEqual(result, tc.expect) {
			t.Logf("result is %v,expect is %v", result, tc.expect)
		}
	}
}
