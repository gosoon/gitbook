package main

import (
	"testing"
)

func TestClimbStairs(t *testing.T) {
	type testCase struct {
		n      int
		expect int
	}

	testCases := []testCase{
		{
			n:      2,
			expect: 2,
		},
		{
			n:      3,
			expect: 3,
		},
	}

	for _, tc := range testCases {
		results := climbStairs(tc.n)
		if results != tc.expect {
			t.Fatalf("result is %v,expect is %v", results, tc.expect)
		}
	}
}
