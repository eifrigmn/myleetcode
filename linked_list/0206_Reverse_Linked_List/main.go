package linked_list

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
