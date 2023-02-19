package tree

func postorder(root *Node) []int {
	var ans []int
	var dfs func(root *Node)
	dfs = func(root *Node) {
		if root == nil {
			return
		}

		for _, child := range root.Children {
			dfs(child)
		}
		ans = append(ans, root.Val)

	}
	dfs(root)
	return ans
}
