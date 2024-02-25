package _0019

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func _removeNthFromEnd(head *ListNode, n int) *ListNode {
	var q = new(ListNode)
	q.Next = head
	var p1 = q
	var p2 = q
	for i := 0; i < n+1; i++ {
		if p1 == nil {
			p1 = head
		} else {
			p1 = p1.Next
		}
	}
	for p1 != nil {
		p1 = p1.Next
		p2 = p2.Next
	}
	p2.Next = p2.Next.Next
	return q.Next
}
