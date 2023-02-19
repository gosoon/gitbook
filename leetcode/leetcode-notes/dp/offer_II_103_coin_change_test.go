package main

import (
	"testing"
)

func TestOfferIICoinChange(t *testing.T) {
	type testCase struct {
		coins  []int
		amount int
		expect int
	}

	testCases := []testCase{
		{
			coins:  []int{1, 2, 5},
			amount: 11,
			expect: 3,
		},
		{
			coins:  []int{2},
			amount: 3,
			expect: -1,
		},
	}

	for _, tc := range testCases {
		results := offerIICoinChange(tc.coins, tc.amount)
		if results != tc.expect {
			t.Fatalf("result is %v,expect is %v", results, tc.expect)
		}
	}
}
