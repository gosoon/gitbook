package tree

/*
// Definition for a Node.
type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	q := []*Node{root}
	newQ := []*Node{}
	for i := 0; i < len(q); i++ {
		if q[i] == nil {
			continue
		}
		if i+1 < len(q) {
			q[i].Next = q[i+1]
		}

		if q[i].Left != nil {
			newQ = append(newQ, q[i].Left)
		}
		if q[i].Right != nil {
			newQ = append(newQ, q[i].Right)
		}

		if i+1 == len(q) {
			i = -1
			q = newQ
			newQ = []*Node{}
		}
	}
	return root
}
*/
