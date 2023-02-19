package main

import (
	"testing"
)

func TestMinWindow(t *testing.T) {
	type testCase struct {
		s      string
		t      string
		expect string
	}

	testCases := []testCase{
		{
			s:      "ADOBECODEBANC",
			t:      "ABC",
			expect: "BANC",
		},
		{
			s:      "a",
			t:      "aa",
			expect: "",
		},
	}

	for _, tc := range testCases {
		result := minWindow(tc.s, tc.t)
		if result != tc.expect {
			t.Logf("result is %v,expect is %v", result, tc.expect)
		}
	}
}
