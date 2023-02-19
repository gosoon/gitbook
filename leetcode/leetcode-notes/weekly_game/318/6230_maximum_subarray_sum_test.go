package main

import (
	"reflect"
	"testing"
)

func TestMaximumSubarraySum(t *testing.T) {
	type testCase struct {
		nums   []int
		k      int
		expect int
	}

	testCases := []testCase{
		//{
		//nums:   []int{1, 5, 4, 2, 9, 9, 9},
		//k:      3,
		//expect: 15,
		//},
		//{
		//nums:   []int{4, 4, 4},
		//k:      3,
		//expect: 0,
		//},
		{
			nums:   []int{1, 1, 1, 7, 8, 9},
			k:      3,
			expect: 24,
		},
	}

	for _, tc := range testCases {
		result := maximumSubarraySum(tc.nums, tc.k)
		t.Logf("result:%v \n", result)
		if !reflect.DeepEqual(result, tc.expect) {
			t.Logf("result is %v,expect is %v", result, tc.expect)
		}
	}

}
