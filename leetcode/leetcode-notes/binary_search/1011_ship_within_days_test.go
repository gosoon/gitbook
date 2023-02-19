package main

import (
	"reflect"
	"testing"
)

func TestShipWithinDays(t *testing.T) {
	type testCase struct {
		nums   []int
		target int
		expect int
	}

	testCases := []testCase{
		{
			nums:   []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			target: 10,
			expect: 10,
		},
		{
			nums:   []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			target: 5,
			expect: 15,
		},
		{
			nums:   []int{3, 2, 2, 4, 1, 4},
			target: 3,
			expect: 6,
		},
		{
			nums:   []int{1, 2, 3, 1, 1},
			target: 4,
			expect: 3,
		},
	}

	for _, tc := range testCases {
		t.Logf("nums:%v,target:%v \n", tc.nums, tc.target)
		result := shipWithinDays(tc.nums, tc.target)
		if !reflect.DeepEqual(result, tc.expect) {
			t.Logf("result is %v,expect is %v", result, tc.expect)
		}
	}
}
