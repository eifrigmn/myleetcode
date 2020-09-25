package _0203

import "myleetcode/util"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode = util.ListNode

func removeElements(head *ListNode, val int) *ListNode {
	prv := &ListNode{}
	prv.Next = head
	l := prv
	for prv.Next != nil {
		if prv.Next.Val == val {
			prv.Next = prv.Next.Next
		} else {
			prv = prv.Next
		}
	}
	return l.Next
}
