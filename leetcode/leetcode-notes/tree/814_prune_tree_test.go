package tree

import (
	"fmt"
	"reflect"
	"testing"
)

func TestPruneTree(t *testing.T) {
	type testCase struct {
		root   []int
		expect []int
	}

	testCases := []testCase{
		{
			root:   []int{1, -1, 0, 0, 1},
			expect: []int{1, 0, 1},
		},
		{
			root:   []int{1, 0, 1, 0, 0, 0, 1},
			expect: []int{1, 1, 1},
		},
	}

	for _, tc := range testCases {
		root := InitTreeByArray(tc.root)
		fmt.Printf("root:%+v \n", PreTraverse(root))
		result := pruneTree(root)
		fmt.Printf("root:%+v \n", PreTraverse(root))
		fmt.Printf("result:%+v \n", PreTraverse(result))
		ans := PreTraverse(result)
		if !reflect.DeepEqual(ans, tc.expect) {
			t.Logf("result is %v,expect is %v", ans, tc.expect)
		}
	}

}
