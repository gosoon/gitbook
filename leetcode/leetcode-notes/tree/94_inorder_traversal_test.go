package tree

import (
	"reflect"
	"testing"
)

func TestInorderTraversal(t *testing.T) {
	type testCase struct {
		root   []int
		expect []int
	}

	testCases := []testCase{
		{
			root:   []int{1, -1, 2, 3},
			expect: []int{1, 3, 2},
		},
	}

	for _, tc := range testCases {
		root := InitTreeByArray(tc.root)
		result := inorderTraversal(root)
		if !reflect.DeepEqual(result, tc.expect) {
			t.Logf("result is %v,expect is %v", result, tc.expect)
		}
	}
}
