package main

import (
	"reflect"
	"testing"
)

func TestGenerateParenthesis(t *testing.T) {
	type testCase struct {
		n      int
		expect []string
	}

	testCases := []testCase{
		{
			n:      3,
			expect: []string{"((()))", "(()())", "(())()", "()(())", "()()()"},
		},
	}

	for _, tc := range testCases {
		results := generateParenthesis(tc.n)
		if !reflect.DeepEqual(tc.expect, results) {
			t.Logf("result is %v,expect is %v", results, tc.expect)
		}
	}
}
