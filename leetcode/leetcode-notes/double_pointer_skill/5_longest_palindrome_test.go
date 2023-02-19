package main

import (
	"reflect"
	"testing"
)

func TestLongestPalindrome(t *testing.T) {
	type testCase struct {
		s      string
		expect string
	}

	testCases := []testCase{
		{
			s:      "babad",
			expect: "bab",
		},
		{
			s:      "cbbd",
			expect: "bb",
		},
		{
			s:      "bb",
			expect: "bb",
		},
	}

	for _, tc := range testCases {
		result := longestPalindrome(tc.s)
		if !reflect.DeepEqual(result, tc.expect) {
			t.Logf("result is %v,expect is %v", result, tc.expect)
		}
	}
}
