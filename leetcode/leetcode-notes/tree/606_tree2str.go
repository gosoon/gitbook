package tree

import (
	"fmt"
	"strconv"
)

/*
func tree2str(root *TreeNode) string {
	var ans string
	var dfs func(root *TreeNode) string
	dfs = func(root *TreeNode) string {
		if root == nil {
			return "()"
		}
		ans += fmt.Sprintf("(%v", root.Val)
		lstr := dfs(root.Left)
		if lstr == "()" && root.Right != nil {
			ans += lstr
		}
		rstr := dfs(root.Right)
		ans = fmt.Sprintf("%s)", ans)
		return ans
	}
	dfs(root)
	return ans[1 : len(ans)-1]
}
*/
func tree2str(root *TreeNode) string {
	if root == nil {
		return ""
	} else if root.Left == nil && root.Right == nil {
		return strconv.Itoa(root.Val)
	} else if root.Right == nil {
		return fmt.Sprintf("%d(%s)", root.Val, tree2str(root.Left))
	} else {
		return fmt.Sprintf("%d(%s)(%s)", root.Val, tree2str(root.Left), tree2str(root.Right))
	}
}
