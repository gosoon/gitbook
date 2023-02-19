package tree

import (
	"fmt"
	"reflect"
	"testing"
)

func TestFindDuplicateSubtrees(t *testing.T) {
	type testCase struct {
		root   []int
		expect [][]int
	}

	testCases := []testCase{
		//{
		//root:   []int{1, 2, 4, -1, 3, 2, 4, -1, 4},
		//expect: [][]int{{2, 4}, {4}},
		//},
		{
			root:   []int{1, 2, 3},
			expect: [][]int{{2, 4}, {4}},
		},
	}

	for _, tc := range testCases {
		root := InitTreeByArray(tc.root)
		fmt.Printf("root:%+v \n", PreTraverse(root))
		result := findDuplicateSubtrees(root)
		if !reflect.DeepEqual(result, tc.expect) {
			t.Logf("result is %v,expect is %v", result, tc.expect)
		}
	}

}
