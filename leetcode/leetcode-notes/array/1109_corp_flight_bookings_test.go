package main

import "testing"

func TestCorpFlightBookings(t *testing.T) {
	type testCase struct {
		bookings [][]int
		n        int
	}

	testCases := []testCase{
		{
			bookings: [][]int{[]int{1, 2, 10}, []int{2, 3, 20}, []int{2, 5, 25}},
			n:        5,
		},
	}

	for _, tc := range testCases {
		t.Log(corpFlightBookings(tc.bookings, tc.n))
	}
}
