package main

import (
	"reflect"
	"testing"
)

func TestPivotInteger(t *testing.T) {
	type testCase struct {
		n      int
		expect int
	}

	testCases := []testCase{
		{
			n:      8,
			expect: 6,
		},
		{
			n:      1,
			expect: 1,
		},
		{
			n:      4,
			expect: -1,
		},
	}

	for _, tc := range testCases {
		results := pivotInteger(tc.n)
		if !reflect.DeepEqual(results, tc.expect) {
			t.Logf("result is %v,expect is %v", results, tc.expect)
		}
	}
}
