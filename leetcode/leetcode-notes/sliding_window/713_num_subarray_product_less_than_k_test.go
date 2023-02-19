package main

import (
	"reflect"
	"testing"
)

func TestNumSubarrayProductLessThanK(t *testing.T) {
	type testCase struct {
		s      []int
		k      int
		expect int
	}

	testCases := []testCase{
		{
			s:      []int{10, 5, 2, 6},
			k:      100,
			expect: 8,
		},
		//{
		//s:      []int{1, 2, 3},
		//k:      0,
		//expect: 0,
		//},
	}

	for _, tc := range testCases {
		t.Log(tc.s)
		result := numSubarrayProductLessThanK(tc.s, tc.k)
		if !reflect.DeepEqual(result, tc.expect) {
			t.Logf("result is %v,expect is %v", result, tc.expect)
		}
	}
}
