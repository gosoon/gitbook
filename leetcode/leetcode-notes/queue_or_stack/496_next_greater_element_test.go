package main

import (
	"reflect"
	"testing"
)

func TestNextGreaterElement(t *testing.T) {
	type testCase struct {
		nums1  []int
		nums2  []int
		expect []int
	}

	testCases := []testCase{
		{
			nums1:  []int{4, 1, 2},
			nums2:  []int{1, 3, 4, 2},
			expect: []int{-1, 3, -1},
		},
	}

	for _, tc := range testCases {
		results := nextGreaterElement(tc.nums1, tc.nums2)
		if !reflect.DeepEqual(results, tc.expect) {
			t.Logf("result is %v,expect is %v", results, tc.expect)
		}
	}
}
