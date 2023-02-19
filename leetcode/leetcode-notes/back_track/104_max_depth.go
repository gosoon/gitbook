package main

func maxDepth(root *TreeNode) int {
	depth := dfs(root, 0)
	return depth
}

func dfs(root *TreeNode, depth int) int {
	if root == nil {
		return depth
	}
	return max(dfs(root.Left, depth+1), dfs(root.Right, depth+1))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
