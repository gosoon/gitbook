package main

func removeNodes(head *ListNode) *ListNode {
	dummy := &ListNode{}

	arr := make([]int, 100005)
	n := 0
	for head != nil {
		arr[n] = head.Val
		head = head.Next
		n++
	}

	last := 0
	for i := n - 1; i >= 0; i-- {
		if arr[i] >= last {
			node := &ListNode{Val: arr[i]}
			node.Next = dummy.Next
			dummy.Next = node
			last = arr[i]
		}
	}
	return dummy.Next
}
