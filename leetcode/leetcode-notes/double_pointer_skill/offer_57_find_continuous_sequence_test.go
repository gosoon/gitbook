package main

import (
	"reflect"
	"testing"
)

func TestFindContinuousSequence(t *testing.T) {
	type testCase struct {
		target int
		expect [][]int
	}

	testCases := []testCase{
		{
			target: 9,
			expect: [][]int{
				[]int{2, 3, 4},
				[]int{4, 5},
			},
		},
		//{
		//target: 15,
		//expect: [][]int{
		//[]int{1, 2, 3, 4, 5},
		//[]int{4, 5, 6},
		//[]int{7, 8},
		//},
		//},
	}

	for _, tc := range testCases {
		result := findContinuousSequence(tc.target)
		if !reflect.DeepEqual(result, tc.expect) {
			t.Logf("result is %v,expect is %v", result, tc.expect)
		}
	}
}
