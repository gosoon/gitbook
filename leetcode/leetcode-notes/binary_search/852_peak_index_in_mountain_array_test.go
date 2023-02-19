package main

import (
	"reflect"
	"testing"
)

func TestPeakIndexInMountainArray(t *testing.T) {
	type testCase struct {
		nums   []int
		expect int
	}

	testCases := []testCase{
		{
			nums:   []int{1, 2, 3, 4, 1},
			expect: 3,
		},
		{
			nums:   []int{0, 1, 0},
			expect: 1,
		},
		{
			nums:   []int{0, 2, 1, 0},
			expect: 1,
		},
		{
			nums:   []int{0, 10, 5, 2},
			expect: 1,
		},
		{
			nums:   []int{3, 4, 5, 1},
			expect: 2,
		},
		{
			nums:   []int{24, 69, 100, 99, 79, 78, 67, 36, 26, 19},
			expect: 2,
		},
	}

	for _, tc := range testCases {
		t.Logf("nums:%v \n", tc.nums)
		result := peakIndexInMountainArray(tc.nums)
		t.Logf("result:%v \n", result)
		if !reflect.DeepEqual(result, tc.expect) {
			t.Logf("result is %v,expect is %v", result, tc.expect)
		}
	}
}
