package main

import (
	"reflect"
	"testing"
)

func TestNextGreaterElements(t *testing.T) {
	type testCase struct {
		nums   []int
		expect []int
	}

	testCases := []testCase{
		{
			nums:   []int{1, 2, 1},
			expect: []int{2, -1, 2},
		},
		{
			nums:   []int{1, 2, 3, 4, 3},
			expect: []int{2, 3, 4, -1, 4},
		},
	}

	for _, tc := range testCases {
		results := nextGreaterElements(tc.nums)
		if !reflect.DeepEqual(results, tc.expect) {
			t.Logf("result is %v,expect is %v", results, tc.expect)
		}
	}
}
