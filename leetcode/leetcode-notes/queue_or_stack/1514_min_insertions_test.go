package main

import (
	"reflect"
	"testing"
)

func TestMinInsertions(t *testing.T) {
	type testCase struct {
		s      string
		expect int
	}

	testCases := []testCase{
		{
			s:      "())",
			expect: 0,
		},
		{
			s:      "(()))",
			expect: 1,
		},
		{
			s:      "))())(",
			expect: 3,
		},
		{
			s:      ")))))))",
			expect: 5,
		},
	}

	for _, tc := range testCases {
		t.Logf(tc.s)
		results := minInsertions(tc.s)
		if !reflect.DeepEqual(results, tc.expect) {
			t.Logf("result is %v,expect is %v", results, tc.expect)
		}
	}
}
