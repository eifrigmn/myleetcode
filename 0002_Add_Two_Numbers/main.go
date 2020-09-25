package _002

import "myleetcode/util"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode = util.ListNode
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var head =new(ListNode)
	var flag int
	h := head
	for l1 != nil {
		val := l1.Val + flag
		if l2 != nil {
			val += l2.Val
			l2 = l2.Next
		}
		h.Next = &ListNode{Val: val%10}
		flag = val/10
		l1 = l1.Next
		h = h.Next
	}
	for l2 != nil {
		h.Next = &ListNode{Val: (l2.Val+flag)%10}
		flag = (l2.Val+flag)/10
		h = h.Next
		l2 = l2.Next
	}
	if flag != 0 {
		h.Next = &ListNode{Val: flag}
	}
	return head.Next
}