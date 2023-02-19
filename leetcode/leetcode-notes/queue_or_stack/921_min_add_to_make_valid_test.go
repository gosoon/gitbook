package main

import (
	"reflect"
	"testing"
)

func TestMinAddToMakeValid(t *testing.T) {
	type testCase struct {
		s      string
		expect int
	}

	testCases := []testCase{
		{
			s:      "())",
			expect: 1,
		},
		{
			s:      "(((",
			expect: 3,
		},
	}

	for _, tc := range testCases {
		results := minAddToMakeValid(tc.s)
		if !reflect.DeepEqual(results, tc.expect) {
			t.Logf("result is %v,expect is %v", results, tc.expect)
		}
	}
}
