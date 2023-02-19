package tree

func pathSum(root *TreeNode, targetSum int) [][]int {
	var ans [][]int
	var res []int
	var traverse func(root *TreeNode, res []int, targetSum int)
	traverse = func(root *TreeNode, res []int, targetSum int) {
		if root == nil {
			return
		}
		res = append(res, root.Val)
		targetSum -= root.Val
		if root.Left == nil && root.Right == nil && targetSum == 0 {
			ans = append(ans, append([]int(nil), res...))
		}
		traverse(root.Left, res, targetSum)
		traverse(root.Right, res, targetSum)
	}
	traverse(root, res, targetSum)
	return ans
}
