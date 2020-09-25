package _0206

import "myleetcode/util"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode = util.ListNode

func reverseList(head *ListNode) *ListNode {
	var rev *ListNode
	for head != nil {
		tmp := head.Next
		head.Next = rev
		rev = head
		head = tmp
	}
	return rev
}

func reverseList1(head *ListNode) *ListNode {
	var prev *ListNode
	for head != nil {
		head.Next, prev, head = prev, head, head.Next
	}
	return prev
}
