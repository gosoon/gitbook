package main

import (
	"reflect"
	"testing"
)

func TestSearchByOffer(t *testing.T) {
	type testCase struct {
		nums   []int
		target int
		expect int
	}

	testCases := []testCase{
		{
			nums:   []int{2, 2},
			target: 3,
			expect: 0,
		},
		{
			nums:   []int{5, 7, 7, 8, 8, 10},
			target: 8,
			expect: 2,
		},
		{
			nums:   []int{5, 7, 7, 8, 8, 10},
			target: 6,
			expect: 0,
		},
	}

	for _, tc := range testCases {
		t.Logf("nums:%v,target:%v \n", tc.nums, tc.target)
		result := searchByOffer(tc.nums, tc.target)
		if !reflect.DeepEqual(result, tc.expect) {
			t.Logf("result is %v,expect is %v", result, tc.expect)
		}
	}
}
