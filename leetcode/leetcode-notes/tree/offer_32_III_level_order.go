package tree

// bfs
func levelOrderOffer32_iii(root *TreeNode) [][]int {
	var ans [][]int
	q := []*TreeNode{root}

	var res []int
	var newq []*TreeNode
	levelOrder := 0
	for i := 0; i < len(q); i++ {
		index := i
		if levelOrder%2 != 0 {
			index = len(q) - 1 - i
		}
		if q[i] != nil {
			if q[i].Left != nil {
				newq = append(newq, q[i].Left)
			}

			if q[i].Right != nil {
				newq = append(newq, q[i].Right)
			}
		}
		if q[index] != nil {
			res = append(res, q[index].Val)
		}

		if i == len(q)-1 {
			i = -1
			q = newq
			newq = []*TreeNode(nil)
			if len(res) != 0 {
				ans = append(ans, res)
			}
			res = []int{}
			levelOrder++
		}
	}
	return ans
}
