package main

import (
	"testing"
)

func TestMinDistance(t *testing.T) {
	type testCase struct {
		s1     string
		s2     string
		expect int
	}

	testCases := []testCase{
		{
			s1:     "horse",
			s2:     "ros",
			expect: 3,
		},
		{
			s1:     "intention",
			s2:     "execution",
			expect: 5,
		},
	}

	for _, tc := range testCases {
		results := minDistance(tc.s1, tc.s2)
		if results != tc.expect {
			t.Fatalf("result is %v,expect is %v", results, tc.expect)
		}
	}
}
