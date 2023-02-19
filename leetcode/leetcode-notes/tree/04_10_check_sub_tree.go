package tree

import (
	"fmt"
	"strings"
)

func checkSubTree(t1 *TreeNode, t2 *TreeNode) bool {
	var dfs func(root *TreeNode) string
	dfs = func(root *TreeNode) string {
		if root == nil {
			return ""
		}

		key := fmt.Sprintf("%v_", root.Val)
		key += dfs(root.Left)
		key += dfs(root.Right)
		return key
	}
	t1Tree := dfs(t1)
	t2Tree := dfs(t2)

	return strings.Contains(t1Tree, t2Tree)
}
