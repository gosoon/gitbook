package main

import (
	"testing"
)

func TestCheckInclusion(t *testing.T) {
	type testCase struct {
		s1     string
		s2     string
		expect bool
	}

	testCases := []testCase{
		{
			s1:     "ab",
			s2:     "eidbaooo",
			expect: true,
		},
		{
			s1:     "ab",
			s2:     "eidboaoo",
			expect: false,
		},
		{
			s1:     "abcdxabcde",
			s2:     "abcdeabcdx",
			expect: true,
		},
		{
			s1:     "hello",
			s2:     "ooolleoooleh",
			expect: false,
		},
	}

	for _, tc := range testCases {
		result := checkInclusion(tc.s1, tc.s2)
		if result != tc.expect {
			t.Logf("result is %v,expect is %v", result, tc.expect)
		}
	}
}
