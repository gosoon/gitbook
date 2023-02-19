package tree

func pruneTree(root *TreeNode) *TreeNode {
	var traverse func(root *TreeNode)
	traverse = func(root *TreeNode) {
		if root == nil {
			return
		}

		traverse(root.Left)
		traverse(root.Right)
		if root.Left != nil && root.Left.Val == 0 && root.Left.Left == nil &&
			root.Left.Right == nil {
			root.Left = nil
		}

		if root.Right != nil && root.Right.Val == 0 && root.Right.Left == nil &&
			root.Right.Right == nil {
			root.Right = nil
		}
	}
	traverse(root)
	if root.Val == 0 && root.Left == nil && root.Right == nil {
		return nil
	}
	return root
}

/*
func pruneTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}

	root.Left = pruneTree(root.Left)
	root.Right = pruneTree(root.Right)
	if root.Val == 0 && root.Left == nil && root.Right == nil {
		return nil
	}
	return root
}
*/
