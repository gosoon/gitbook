package tree

func isUnivalTree(root *TreeNode) bool {
	if root == nil {
		return true
	}
	val := root.Val

	var traverse func(root *TreeNode) bool
	traverse = func(root *TreeNode) bool {
		if root == nil {
			return true
		}
		if root.Val != val {
			return false
		}
		isLeftTrue := traverse(root.Left)
		isRightTrue := traverse(root.Right)
		return isLeftTrue && isRightTrue
	}
	return traverse(root)
}
