package _0083

import "myleetcode/util"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode = util.ListNode

func deleteDuplicates(head *ListNode) *ListNode {
	result := head
	if head == nil || head.Next == nil {
		return head
	}
	for head.Next != nil {
		if head.Val == head.Next.Val {
			head.Next = head.Next.Next
		} else {
			head = head.Next
		}
	}
	return result
}

func deleteDuplicates1(head *ListNode) *ListNode {
	var h = head
	for h != nil {
		if h.Next != nil && h.Val == h.Next.Val {
			h.Next = h.Next.Next
			continue
		}
		h = h.Next
	}
	return head
}
