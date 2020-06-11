package _0143

import "myleetcode/util"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode = util.ListNode

func reorderList(head *ListNode) {
	if head == nil || head.Next == nil {
		return
	}
	mid := findMiddle(head)
	rgt := mid.Next
	rgt = reverse(rgt)
	mid.Next = nil
	curr := head
	for curr != nil && rgt != nil {
		tmpRgt := rgt.Next
		tmpCurr := curr.Next
		curr.Next = rgt
		rgt.Next = tmpCurr
		rgt = tmpRgt
		curr = tmpCurr
	}
}

// 找链表的中点
func findMiddle(head *ListNode) *ListNode {
	fast, slow := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	return slow
}

func reverse(head *ListNode) *ListNode {
	var prev *ListNode
	for head != nil {
		head.Next, prev, head = prev, head, head.Next
	}
	return prev
}
