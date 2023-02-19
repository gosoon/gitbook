package main

import (
	"reflect"
	"testing"
)

func TestMinEatingSpeed(t *testing.T) {
	type testCase struct {
		nums   []int
		target int
		expect int
	}

	testCases := []testCase{
		{
			nums:   []int{2, 2},
			target: 2,
			expect: 2,
		},
		{
			nums:   []int{3, 6, 7, 11},
			target: 8,
			expect: 4,
		},
		{
			nums:   []int{30, 11, 23, 4, 20},
			target: 5,
			expect: 30,
		},
		{
			nums:   []int{30, 11, 23, 4, 20},
			target: 6,
			expect: 23,
		},
	}

	for _, tc := range testCases {
		t.Logf("nums:%v,target:%v \n", tc.nums, tc.target)
		result := minEatingSpeed(tc.nums, tc.target)
		if !reflect.DeepEqual(result, tc.expect) {
			t.Logf("result is %v,expect is %v", result, tc.expect)
		}
	}
}
