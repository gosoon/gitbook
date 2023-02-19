package main

import (
	"reflect"
	"testing"
)

func TestLetterCombinations(t *testing.T) {
	type testCase struct {
		digits string
		expect []string
	}

	testCases := []testCase{
		{
			digits: "23",
			expect: []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"},
		},
		{
			digits: "2",
			expect: []string{"a", "b", "c"},
		},
		{
			digits: "",
			expect: []string{},
		},
	}

	for _, tc := range testCases {
		results := letterCombinations(tc.digits)
		if !reflect.DeepEqual(tc.expect, results) {
			t.Logf("result is %v,expect is %v", results, tc.expect)
		}
	}
}
