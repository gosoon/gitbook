package main

import (
	"reflect"
	"testing"
)

func TestSortPeople(t *testing.T) {
	type testCase struct {
		names   []string
		heights []int
		expect  []string
	}

	testCases := []testCase{
		{
			names:   []string{"Mary", "John", "Emma"},
			heights: []int{180, 165, 170},
			expect:  []string{"Mary", "Emma", "John"},
		},
	}

	for _, tc := range testCases {
		result := sortPeople(tc.names, tc.heights)
		t.Logf("result:%v \n", result)
		if !reflect.DeepEqual(result, tc.expect) {
			t.Logf("result is %v,expect is %v", result, tc.expect)
		}
	}

}
