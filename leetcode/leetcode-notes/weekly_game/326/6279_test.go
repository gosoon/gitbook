package main

import (
	"reflect"
	"testing"
)

func TestDistinctPrimeFactors(t *testing.T) {
	type testCase struct {
		nums   []int
		expect int
	}

	testCases := []testCase{
		{
			nums:   []int{2, 4, 3, 7, 10, 6},
			expect: 4,
		},
		{
			nums:   []int{2, 4, 8, 16},
			expect: 1,
		},
		{
			nums:   []int{20, 16, 5, 19, 7, 13, 8, 8},
			expect: 5,
		},
	}

	for _, tc := range testCases {
		results := distinctPrimeFactors(tc.nums)
		if !reflect.DeepEqual(results, tc.expect) {
			t.Fatalf("result is %v,expect is %v", results, tc.expect)
		}
	}
}
