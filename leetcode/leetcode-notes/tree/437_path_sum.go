package tree

func pathSum437(root *TreeNode, targetSum int) int {
	var ans int
	var traverse func(root *TreeNode, targetSum int)
	traverse = func(root *TreeNode, targetSum int) {
		if root == nil {
			return
		}
		traverse(root.Left, targetSum-root.Val)
		traverse(root.Right, targetSum-root.Val)
		if targetSum-root.Val == 0 {
			ans++
		}
	}
	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		traverse(root, targetSum)
		dfs(root.Left)
		dfs(root.Right)
	}
	dfs(root)
	return ans
}
