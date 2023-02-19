package main

import (
	"reflect"
	"testing"
)

func TestFindAnagrams(t *testing.T) {
	type testCase struct {
		s      string
		p      string
		expect []int
	}

	testCases := []testCase{
		{
			s:      "cbaebabacd",
			p:      "abc",
			expect: []int{0, 6},
		},
		{
			s:      "abab",
			p:      "ab",
			expect: []int{0, 1, 2},
		},
		{
			s:      "baa",
			p:      "aa",
			expect: []int{1},
		},
	}

	for _, tc := range testCases {
		result := findAnagrams(tc.s, tc.p)
		if !reflect.DeepEqual(result, tc.expect) {
			t.Logf("result is %v,expect is %v", result, tc.expect)
		}
	}
}
