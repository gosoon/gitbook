package main

import "testing"

func TestCheckDistances(t *testing.T) {
	type testCase struct {
		s        string
		distance []int
	}

	testCases := []testCase{
		{
			s:        "abaccb",
			distance: []int{1, 3, 0, 5, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		},
		{
			s:        "aa",
			distance: []int{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		},
		{
			s:        "zz",
			distance: []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
		},
	}

	for _, tc := range testCases {
		t.Log(checkDistances(tc.s, tc.distance))
	}
}
