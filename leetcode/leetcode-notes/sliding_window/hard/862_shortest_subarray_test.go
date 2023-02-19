package main

import (
	"reflect"
	"testing"
)

func TestShortestSubarray(t *testing.T) {
	type testCase struct {
		nums   []int
		k      int
		expect int
	}

	testCases := []testCase{
		{
			nums:   []int{17, 85, 93, -45, -21},
			k:      150,
			expect: 2,
		},
		{
			nums:   []int{48, 99, 37, 4, -31},
			k:      140,
			expect: 2,
		},
		{
			nums:   []int{1},
			k:      1,
			expect: 1,
		},
		{
			nums:   []int{1, 2},
			k:      4,
			expect: -1,
		},
		{
			nums:   []int{2, -1, 2},
			k:      3,
			expect: 3,
		},
	}

	for _, tc := range testCases {
		t.Log(tc.nums)
		result := shortestSubarray(tc.nums, tc.k)
		if !reflect.DeepEqual(result, tc.expect) {
			t.Logf("result is %v,expect is %v", result, tc.expect)
		}
	}
}
