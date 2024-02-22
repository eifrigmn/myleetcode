package _0083

import "myleetcode/util"

type ListNode = util.ListNode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func partition(head *ListNode, x int) *ListNode {
	var lower, bigger = new(ListNode), new(ListNode)
	var ll = lower
	var lb = bigger
	for head != nil {
		if head.Val >= x {
			lb.Next = &ListNode{
				Val: head.Val,
			}
			lb = lb.Next
		} else {
			ll.Next = &ListNode{
				Val: head.Val,
			}
			ll = ll.Next
		}
		head = head.Next
	}
	ll.Next = bigger.Next
	return lower.Next
}
