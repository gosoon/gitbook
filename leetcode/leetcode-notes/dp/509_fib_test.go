package main

import (
	"testing"
)

func TestFib(t *testing.T) {
	type testCase struct {
		n      int
		expect int
	}

	testCases := []testCase{
		{
			n:      2,
			expect: 1,
		},
		{
			n:      4,
			expect: 3,
		},
	}

	for _, tc := range testCases {
		results := fib(tc.n)
		if results != tc.expect {
			t.Logf("result is %v,expect is %v", results, tc.expect)
		}
	}
}
