package tree

import (
	"fmt"

	"github.com/gosoon/algorithms-golang/data-structures/queue"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// init tree
func InitTreeByArray(nums []int) *TreeNode {
	q := queue.NewArrayQueue(len(nums))
	for _, v := range nums {
		q.Enqueue(v)
	}

	root := InitTree(q)
	return root
}

func InitTree(q *queue.ArrayQueue) *TreeNode {
	v, ok := q.Dequeue()
	if !ok {
		return nil
	}
	if v.(int) == -1 {
		return nil
	}
	root := &TreeNode{Val: v.(int)}
	root.Left = InitTree(q)
	root.Right = InitTree(q)
	return root
}

func PreTraverse(root *TreeNode) []int {
	var ans []int
	var traverse func(root *TreeNode)
	traverse = func(root *TreeNode) {
		if root == nil {
			return
		}
		ans = append(ans, root.Val)
		traverse(root.Left)
		traverse(root.Right)
	}
	traverse(root)
	return ans
}

func InsertNodeToTree(tree *TreeNode, node *TreeNode) {
	if tree == nil || node == nil {
		return
	}

	if tree.Val == 0 {
		tree.Val = node.Val
		return
	}

	if node.Val > tree.Val {
		if tree.Right == nil {
			tree.Right = &TreeNode{}
		}
		InsertNodeToTree(tree.Right, node)
	}
	if node.Val < tree.Val {
		if tree.Left == nil {
			tree.Left = &TreeNode{}
		}
		InsertNodeToTree(tree.Left, node)
	}
}

func InitTree2(Vals []int) (root *TreeNode) {
	rootNode := &TreeNode{Val: Vals[0]}
	for _, val := range Vals {
		var node *TreeNode
		if val != -1 {
			node = &TreeNode{Val: val}
		}
		InsertNodeToTree(rootNode, node)
	}
	return rootNode
}

func MidPrint(root *TreeNode) {
	if root == nil {
		return
	}
	//MidPrint(root.Left)
	fmt.Printf("%v  ", root.Val)
	MidPrint(root.Right)
}
