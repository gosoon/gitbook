package main

import (
	"reflect"
	"testing"
)

func TestMinSubArrayLen(t *testing.T) {
	type testCase struct {
		nums   []int
		target int
		expect int
	}

	testCases := []testCase{
		{
			nums:   []int{2, 3, 1, 2, 4, 3},
			target: 7,
			expect: 2,
		},
		{
			nums:   []int{1, 4, 4},
			target: 4,
			expect: 1,
		},
		{
			nums:   []int{1, 1, 1, 1, 1, 1, 1, 1},
			target: 11,
			expect: 0,
		},
		{
			nums:   []int{1, 2, 3, 4, 5},
			target: 11,
			expect: 3,
		},
	}

	for _, tc := range testCases {
		result := minSubArrayLen(tc.target, tc.nums)
		if !reflect.DeepEqual(result, tc.expect) {
			t.Logf("result is %v,expect is %v", result, tc.expect)
		}
	}
}
