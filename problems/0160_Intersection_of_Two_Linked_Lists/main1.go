package _0160

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func _getIntersectionNode(headA, headB *ListNode) *ListNode {
	var p1 = headA
	var p2 = headB
	for p1 != p2 {
		if p1 == nil {
			p1 = headB
		}
		if p2 == nil {
			p2 = headA
		}
		p1 = p1.Next
		p2 = p2.Next
	}
	return p1
}
