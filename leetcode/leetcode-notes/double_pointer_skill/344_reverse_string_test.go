package main

import (
	"reflect"
	"testing"
)

func TestReverseString(t *testing.T) {
	type testCase struct {
		s      []byte
		expect []byte
	}

	testCases := []testCase{
		{
			s:      []byte{'h', 'e', 'l', 'l', 'o'},
			expect: []byte{'o', 'l', 'l', 'e', 'h'},
		},
		{
			s:      []byte{'H', 'a', 'n', 'n', 'a', 'h'},
			expect: []byte{'h', 'a', 'n', 'n', 'a', 'H'},
		},
	}

	for _, tc := range testCases {
		reverseString(tc.s)
		if !reflect.DeepEqual(tc.s, tc.expect) {
			t.Logf("result is %v,expect is %v", tc.s, tc.expect)
		}
	}
}
