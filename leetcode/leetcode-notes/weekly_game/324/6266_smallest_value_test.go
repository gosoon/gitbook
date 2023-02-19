package main

import (
	"reflect"
	"testing"
)

func TestSmallestValue(t *testing.T) {
	type testCase struct {
		n      int
		expect int
	}

	testCases := []testCase{
		{
			n:      4,
			expect: 4,
		},
		{
			n:      12,
			expect: 7,
		},
		{
			n:      15,
			expect: 5,
		},
		{
			n:      3,
			expect: 3,
		},
	}

	for _, tc := range testCases {
		results := smallestValue(tc.n)
		if !reflect.DeepEqual(results, tc.expect) {
			t.Fatalf("result is %v,expect is %v", results, tc.expect)
		}
	}
}
