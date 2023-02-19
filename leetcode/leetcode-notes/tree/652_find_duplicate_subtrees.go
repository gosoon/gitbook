package tree

import "fmt"

func findDuplicateSubtrees(root *TreeNode) []*TreeNode {
	ans := make([]*TreeNode, 0)
	m := make(map[string]int)
	var dfs func(root *TreeNode) string
	dfs = func(root *TreeNode) string {
		if root == nil {
			return ""
		}
		key := fmt.Sprintf("%v_", root.Val)
		key += dfs(root.Left)
		key += dfs(root.Right)

		fmt.Printf("key:%v \n", key)
		m[key]++
		if count := m[key]; count == 2 {
			ans = append(ans, root)
		}
		return key
	}

	dfs(root)
	return ans
}
