package main

import (
	"testing"
)

func TestLengthOfLongestSubstring(t *testing.T) {
	type testCase struct {
		s      string
		expect int
	}

	testCases := []testCase{
		//{
		//s:      "abcabcbb",
		//expect: 3,
		//},
		//{
		//s:      "bbbbb",
		//expect: 1,
		//},
		//{
		//s:      "pwwkew",
		//expect: 3,
		//},
		{
			s:      "aab",
			expect: 2,
		},
	}

	for _, tc := range testCases {
		result := lengthOfLongestSubstring(tc.s)
		if result != tc.expect {
			t.Logf("result is %v,expect is %v", result, tc.expect)
		}
	}
}
