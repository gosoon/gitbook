package tree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

/*
func preorderTraversal(root *TreeNode) []int {
	var ans []int
	var traversal func(root *TreeNode)
	traversal = func(root *TreeNode) {
		if root == nil {
			return
		}
		ans = append(ans, root.Val)
		traversal(root.Left)
		traversal(root.Right)
	}
	traversal(root)
	return res
}
*/
func preorderTraversal(root *TreeNode) []int {
	var ans []int
	if root == nil {
		return ans
	}
	ans = append(ans, root.Val)
	ans = append(ans, preorderTraversal(root.Left)...)
	ans = append(ans, preorderTraversal(root.Right)...)
	return ans
}
