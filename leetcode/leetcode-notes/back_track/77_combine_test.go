package main

import (
	"testing"
)

func TestCombine(t *testing.T) {
	type testCase struct {
		n      int
		k      int
		expect [][]int
	}

	testCases := []testCase{
		{
			n: 4,
			k: 2,
		},
		{
			n: 1,
			k: 1,
		},
	}

	for _, tc := range testCases {
		results := combine(tc.n, tc.k)
		t.Logf("result is %v,expect is %v", results, tc.expect)
	}
}
