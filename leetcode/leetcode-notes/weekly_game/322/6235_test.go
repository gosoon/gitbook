package main

import (
	"reflect"
	"testing"
)

func TestIsCircularSentence(t *testing.T) {
	type testCase struct {
		sentence string
		expect   bool
	}

	testCases := []testCase{
		{
			sentence: "leetcode exercises sound delightful",
			expect:   true,
		},
		{
			sentence: "eetcode",
			expect:   true,
		},
	}

	for _, tc := range testCases {
		results := isCircularSentence(tc.sentence)
		if !reflect.DeepEqual(results, tc.expect) {
			t.Logf("result is %v,expect is %v", results, tc.expect)
		}
	}
}
