package main

func reversePrint(head *ListNode) []int {
	if head == nil {
		return nil
	}
	res := reversePrint(head.Next)
	return append(res, head.Val)
}

/*
func reversePrint(head *ListNode) []int {
	var ans []int
	var traverse func(head *ListNode)
	traverse = func(head *ListNode) {
		if head == nil {
			return
		}
		traverse(head.Next)
		ans = append(ans, head.Val)
	}
	traverse(head)
	return ans
}
*/
