package _04

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists1(l1 *ListNode, l2 *ListNode) *ListNode {
	var l = new(ListNode)
	var head = l
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			head.Next = l1
			l1 = l1.Next
		} else if l1.Val > l2.Val {
			head.Next = l2
			l2 = l2.Next
		} else {
			l1t := l1.Next
			l2t := l2.Next
			head.Next = l1
			head = head.Next
			head.Next = l2
			l1 = l1t
			l2 = l2t
		}
		head = head.Next
	}

	if l1 != nil {
		head.Next = l1
	}

	if l2 != nil {
		head.Next = l2
	}
	return l.Next
}
