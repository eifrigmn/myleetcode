package _04

import "myleetcode/util"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode = util.ListNode
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	l1 = &ListNode{Next: l1}
	var l = l1
	for l2 != nil && l1.Next != nil {
		if l1.Next.Val >= l2.Val {
			tmp := l2.Next
			l2.Next = l1.Next
			l1.Next = l2
			l2 = tmp
		}
		l1 = l1.Next
	}

	if l2 != nil {
		l1.Next = l2
	}
	return l.Next
}
