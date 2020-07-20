package _02

import "myleetcode/util"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode = util.ListNode
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	var h, p = head, head
	for h != nil {
		if n < 0 {
			p = p.Next
		}
		h = h.Next
		n--
	}

	if n == 0 {
		return head.Next
	}
	p.Next = p.Next.Next
	return head
}
