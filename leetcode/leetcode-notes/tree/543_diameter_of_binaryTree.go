package tree

func diameterOfBinaryTree(root *TreeNode) int {
	var ans int
	var dfs func(root *TreeNode) int
	dfs = func(root *TreeNode) int {
		if root == nil {
			return 0
		}

		leftDepth := dfs(root.Left)
		rightDepth := dfs(root.Right)
		if leftDepth+rightDepth > ans {
			ans = leftDepth + rightDepth
		}

		if leftDepth > rightDepth {
			return leftDepth + 1
		}

		return rightDepth + 1
		//return max(leftDepth, rightDepth) +1
	}
	return dfs(root)
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
