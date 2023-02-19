package main

import (
	"reflect"
	"testing"
)

func TestCountDigits(t *testing.T) {
	type testCase struct {
		n      int
		expect int
	}

	testCases := []testCase{
		{
			n:      7,
			expect: 1,
		},
		{
			n:      121,
			expect: 2,
		},
		{
			n:      1248,
			expect: 4,
		},
	}

	for _, tc := range testCases {
		results := countDigits(tc.n)
		if !reflect.DeepEqual(results, tc.expect) {
			t.Logf("result is %v,expect is %v", results, tc.expect)
		}
	}
}
