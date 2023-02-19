package main

import (
	"reflect"
	"testing"
)

func TestCanPartitionKSubsets(t *testing.T) {
	type testCase struct {
		nums   []int
		k      int
		expect bool
	}

	testCases := []testCase{
		//{
		//nums:   []int{4, 3, 2, 3, 5, 2, 1},
		//k:      4,
		//expect: true,
		//},
		//{
		//nums:   []int{1, 2, 3, 4},
		//k:      3,
		//expect: false,
		//},
		{
			nums:   []int{3, 3, 10, 2, 6, 5, 10, 6, 8, 3, 2, 1, 6, 10, 7, 2},
			k:      6,
			expect: true,
		},
	}

	for _, tc := range testCases {
		results := canPartitionKSubsets(tc.nums, tc.k)
		if !reflect.DeepEqual(results, tc.expect) {
			t.Logf("result is %v,expect is %v", results, tc.expect)
		}
	}
}
