package main

import (
	"reflect"
	"testing"
)

func TestNthUglyNumber(t *testing.T) {
	type testCase struct {
		n      int
		expect int
	}

	testCases := []testCase{
		{
			n:      1,
			expect: 1,
		},
		{
			n:      10,
			expect: 12,
		},
		{
			n:      1352,
			expect: 12,
		},
	}

	for _, tc := range testCases {
		result := nthUglyNumber(tc.n)
		if !reflect.DeepEqual(result, tc.expect) {
			t.Logf("result is %v,expect is %v", result, tc.expect)
		}
	}
}
