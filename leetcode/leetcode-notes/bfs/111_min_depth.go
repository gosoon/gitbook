package main

// bfs
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	q := []*TreeNode{root}

	depth := 1
	var newq []*TreeNode
	for i := 0; i < len(q); i++ {
		if q[i] != nil {
			if q[i].Left == nil && q[i].Right == nil {
				return depth
			}

			if q[i].Left != nil {
				newq = append(newq, q[i].Left)
			}

			if q[i].Right != nil {
				newq = append(newq, q[i].Right)
			}
		}
		if i == len(q)-1 {
			depth++
			i = -1
			q = newq
			newq = []*TreeNode{}
		}
	}
	return depth
}
