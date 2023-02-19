package main

import (
	"testing"
)

func TestCombinationSum3(t *testing.T) {
	type testCase struct {
		k      int
		n      int
		expect [][]int
	}

	testCases := []testCase{
		{
			k:      3,
			n:      7,
			expect: [][]int{},
		},
		{
			k:      3,
			n:      9,
			expect: [][]int{},
		},
		{
			k:      4,
			n:      1,
			expect: [][]int{},
		},
	}

	for _, tc := range testCases {
		results := combinationSum3(tc.k, tc.n)
		t.Logf("result is %v,expect is %v", results, tc.expect)
	}
}
