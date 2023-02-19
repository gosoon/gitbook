package main

import (
	"testing"
)

func TestSumOfNumberAndReverse(t *testing.T) {
	type testCase struct {
		nums   int
		expect bool
	}

	testCases := []testCase{
		{
			nums:   443,
			expect: true,
		},
		{
			nums:   181,
			expect: true,
		},
		{
			nums:   0,
			expect: true,
		},
	}

	for _, tc := range testCases {
		results := sumOfNumberAndReverse(tc.nums)
		if results != tc.expect {
			t.Logf("result is %v,expect is %v", results, tc.expect)
		}
	}
}
