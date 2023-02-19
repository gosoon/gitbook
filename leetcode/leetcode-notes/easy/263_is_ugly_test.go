package main

import (
	"reflect"
	"testing"
)

func TestIsUgly(t *testing.T) {
	type testCase struct {
		n      int
		expect bool
	}

	testCases := []testCase{
		{
			n:      6,
			expect: true,
		},
		{
			n:      8,
			expect: true,
		},
		{
			n:      0,
			expect: false,
		},
	}

	for _, tc := range testCases {
		result := isUgly(tc.n)
		if !reflect.DeepEqual(result, tc.expect) {
			t.Logf("result is %v,expect is %v", result, tc.expect)
		}
	}
}
