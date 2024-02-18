package _002

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func _addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var t int
	var result = new(ListNode)
	tail := result
	for l1 != nil {
		var val int
		if l2 == nil {
			val = t + l1.Val
		} else {
			val = t + l1.Val + l2.Val
			l2 = l2.Next
		}
		tail.Next = &ListNode{
			Val:  val % 10,
			Next: nil,
		}
		t = val / 10
		tail = tail.Next
		l1 = l1.Next
	}
	for l2 != nil {
		val := t + l2.Val
		tail.Next = &ListNode{
			Val:  val % 10,
			Next: nil,
		}
		t = val / 10
		tail = tail.Next
		l2 = l2.Next
	}
	if t > 0 {
		tail.Next = &ListNode{
			Val:  t,
			Next: nil,
		}
	}
	return result.Next
}
