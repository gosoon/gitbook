package tree

// Definition for a Node.
type Node struct {
	Val      int
	Children []*Node
}

/*
func preorder(root *Node) []int {
	var ans []int
	if root == nil {
		return ans
	}

	ans = append(ans, root.Val)
	for _, child := range root.Children {
		ans = append(ans, preorder(child)...)
	}
	return ans
}
*/
func preorder(root *Node) []int {
	var ans []int
	var traverse func(root *Node)
	traverse = func(root *Node) {
		if root == nil {
			return
		}

		ans = append(ans, root.Val)
		for _, child := range root.Children {
			traverse(child)
		}
	}
	traverse(root)
	return ans
}
