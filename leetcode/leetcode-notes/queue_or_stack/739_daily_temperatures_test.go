package main

import (
	"reflect"
	"testing"
)

func TestDailyTemperatures(t *testing.T) {
	type testCase struct {
		nums   []int
		expect []int
	}

	testCases := []testCase{
		{
			nums:   []int{73, 74, 75, 71, 69, 72, 76, 73},
			expect: []int{1, 1, 4, 2, 1, 1, 0, 0},
		},
		{
			nums:   []int{30, 40, 50, 60},
			expect: []int{1, 1, 1, 0},
		},
		{
			nums:   []int{30, 60, 90},
			expect: []int{1, 1, 0},
		},
	}

	for _, tc := range testCases {
		results := dailyTemperatures(tc.nums)
		if !reflect.DeepEqual(results, tc.expect) {
			t.Logf("result is %v,expect is %v", results, tc.expect)
		}
	}
}
