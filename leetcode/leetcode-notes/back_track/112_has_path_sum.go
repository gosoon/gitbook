package main

//  Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}

	if root.Left == nil && root.Right == nil {
		return root.Val == targetSum
	}
	return hasPathSum(root.Left, targetSum-root.Val) || hasPathSum(root.Right, targetSum-root.Val)
}

/*
func backTrack(root *TreeNode, sum int, targetSum int) bool {
	fmt.Println(root.Val, sum)
	if root.Left == nil && root.Right == nil {
		sum += root.Val
		fmt.Printf("---> root.Val:%v sum:%v \n", root.Val, sum)
		if sum == targetSum {
			return true
		}
	}

	if root.Left != nil {
		if backTrack(root.Left, sum+root.Val, targetSum) == true {
			return true
		}
	}

	if root.Right != nil {
		if backTrack(root.Right, sum+root.Val, targetSum) == true {
			return true
		}
	}
	return false
}
*/
