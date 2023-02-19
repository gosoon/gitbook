package main

import (
	"reflect"
	"testing"
)

func TestPermutation(t *testing.T) {
	type testCase struct {
		s      string
		expect []string
	}

	testCases := []testCase{
		{
			s:      "abc",
			expect: []string{"abc", "acb", "bac", "bca", "cab", "cba"},
		},
		{
			s:      "aab",
			expect: []string{"aba", "aab", "baa"},
		},
	}

	for _, tc := range testCases {
		results := permutation(tc.s)
		if !reflect.DeepEqual(tc.expect, results) {
			t.Logf("result is %v,expect is %v", results, tc.expect)
		}
	}
}
