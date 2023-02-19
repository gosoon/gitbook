package main

import (
	"reflect"
	"testing"
)

func TestSearch(t *testing.T) {
	type testCase struct {
		nums   []int
		target int
		expect int
	}

	testCases := []testCase{
		{
			nums:   []int{-1, 0, 3, 5, 9, 12},
			target: 9,
			expect: 4,
		},
		{
			nums:   []int{-1, 0, 3, 5, 9, 12},
			target: 2,
			expect: -1,
		},
	}

	for _, tc := range testCases {
		t.Log(tc.nums, tc.target)
		result := search(tc.nums, tc.target)
		if !reflect.DeepEqual(result, tc.expect) {
			t.Logf("result is %v,expect is %v", result, tc.expect)
		}
	}
}
