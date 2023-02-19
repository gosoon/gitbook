package main

import (
	"reflect"
	"testing"
)

func TestContainsNearbyAlmostDuplicate(t *testing.T) {
	type testCase struct {
		nums   []int
		k      int
		t      int
		expect bool
	}

	testCases := []testCase{
		{
			nums:   []int{7, 2, 8},
			k:      2,
			t:      1,
			expect: true,
		},
		{
			nums:   []int{2, 2},
			k:      2,
			t:      0,
			expect: true,
		},
		{
			nums:   []int{1, 2, 3, 1},
			k:      3,
			t:      0,
			expect: true,
		},
		{
			nums:   []int{1, 0, 1, 1},
			k:      1,
			t:      2,
			expect: true,
		},
		{
			nums:   []int{1, 5, 9, 1, 5, 9},
			k:      2,
			t:      3,
			expect: false,
		},
		{
			nums:   []int{7, 1, 3},
			k:      2,
			t:      3,
			expect: true,
		},
		{
			nums:   []int{1, 2, 2, 3, 4, 5},
			k:      3,
			t:      0,
			expect: true,
		},
	}

	for _, tc := range testCases {
		t.Log(tc.nums)
		result := containsNearbyAlmostDuplicate(tc.nums, tc.k, tc.t)
		if !reflect.DeepEqual(result, tc.expect) {
			t.Logf("result is %v,expect is %v", result, tc.expect)
		}
	}
}
