package main

import (
	"reflect"
	"testing"
)

func TestRomanToInt(t *testing.T) {
	type testCase struct {
		s      string
		expect int
	}

	testCases := []testCase{
		{
			s:      "III",
			expect: 3,
		},
		{
			s:      "IV",
			expect: 4,
		},
		{
			s:      "MCMXCIV",
			expect: 1994,
		},
	}

	for _, tc := range testCases {
		result := romanToInt(tc.s)
		if !reflect.DeepEqual(result, tc.expect) {
			t.Logf("result is %v,expect is %v", result, tc.expect)
		}
	}
}
