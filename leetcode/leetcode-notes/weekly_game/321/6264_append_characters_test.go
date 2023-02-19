package main

import (
	"reflect"
	"testing"
)

func TestAppendCharacters(t *testing.T) {
	type testCase struct {
		s      string
		t      string
		expect int
	}

	testCases := []testCase{
		{
			s:      "coaching",
			t:      "coding",
			expect: 4,
		},
	}

	for _, tc := range testCases {
		results := appendCharacters(tc.s, tc.t)
		if !reflect.DeepEqual(results, tc.expect) {
			t.Logf("result is %v,expect is %v", results, tc.expect)
		}
	}
}
