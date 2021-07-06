package _03

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList2(head *ListNode) *ListNode {
	var newHead = new(ListNode)
	for head != nil {
		tmp := head.Next
		head.Next = newHead.Next
		newHead.Next = head
		head = tmp
	}
	return newHead.Next
}
