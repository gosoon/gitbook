package tree

// bfs
func levelOrder(root *TreeNode) [][]int {
	var ans [][]int
	q := []*TreeNode{root}

	var res []int
	var newq []*TreeNode
	for i := 0; i < len(q); i++ {
		if q[i] != nil {
			res = append(res, q[i].Val)
			if q[i].Left != nil {
				newq = append(newq, q[i].Left)
			}

			if q[i].Right != nil {
				newq = append(newq, q[i].Right)
			}
		}
		if i == len(q)-1 {
			i = -1
			q = newq
			newq = []*TreeNode{}
			if len(res) != 0 {
				ans = append(ans, res)
			}
			res = []int{}
		}
	}
	return ans
}
