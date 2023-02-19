package tree

func maxDepth(root *Node) int {
	var level int
	if root == nil {
		return level
	}
	var dfs func(root *Node, depth int)
	dfs = func(root *Node, depth int) {
		if root == nil {
			return
		}
		if depth > level {
			level = depth
		}
		for _, child := range root.Children {
			dfs(child, depth+1)
		}
	}
	dfs(root, level+1)
	return level
}
