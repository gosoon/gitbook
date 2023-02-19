package main

import (
	"reflect"
	"testing"
)

func TestSimilarPairs(t *testing.T) {
	type testCase struct {
		words  []string
		expect int
	}

	testCases := []testCase{
		{
			words:  []string{"aba", "aabb", "abcd", "bac", "aabc"},
			expect: 2,
		},
		{
			words:  []string{"aabb", "ab", "ba"},
			expect: 3,
		},
	}

	for _, tc := range testCases {
		results := similarPairs(tc.words)
		if !reflect.DeepEqual(results, tc.expect) {
			t.Fatalf("result is %v,expect is %v", results, tc.expect)
		}
	}
}
