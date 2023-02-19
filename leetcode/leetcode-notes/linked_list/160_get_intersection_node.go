package main

import "fmt"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	preA := headA
	preB := headB
	var lenA, lenB int
	for preA != nil {
		preA = preA.Next
		lenA++
	}
	for preB != nil {
		preB = preB.Next
		lenB++
	}
	fmt.Println(lenA, lenB)

	if lenA > lenB {
		diff := lenA - lenB
		for i := 1; i <= diff; i++ {
			headA = headA.Next
		}
	} else if lenA < lenB {
		diff := lenB - lenA
		for i := 1; i <= diff; i++ {
			headB = headB.Next
		}
	}

	for headA != headB {
		headA = headA.Next
		headB = headB.Next
	}
	return headA
}
