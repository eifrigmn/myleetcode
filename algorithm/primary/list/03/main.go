package _03

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
	var l =&ListNode{}
	for head != nil {
		t := head.Next
		head.Next = l.Next
		l.Next = head
		head = t
	}
	return l.Next
}

// 递归方法
// 找到原始链表的最后一个节点并返回
func reverseList1(head *ListNode) *ListNode{
	if head == nil || head.Next == nil{
		return head
	}
	newHead := reverseList1(head.Next)
	head.Next.Next = head
	head.Next = nil
	return newHead
}