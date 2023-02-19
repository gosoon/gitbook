package main

import (
	"reflect"
	"testing"
)

func TestGetDistances(t *testing.T) {
	type testCase struct {
		nums   []int
		expect []int
	}

	testCases := []testCase{
		{
			nums:   []int{2, 1, 3, 1, 2, 3, 3},
			expect: []int{4, 2, 7, 2, 4, 4, 5},
		},
		{
			nums:   []int{10, 5, 10, 10},
			expect: []int{5, 0, 3, 4},
		},
	}

	for _, tc := range testCases {
		t.Logf("nums is %+v \n", tc.nums)
		result := getDistances(tc.nums)
		if !reflect.DeepEqual(result, tc.expect) {
			t.Logf("result is %+v,expect is %+v", result, tc.expect)
		}
	}
}
