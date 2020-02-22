package linked_list

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	// 设立观察哨
	r := &ListNode{}
	list := r
	for l1 != nil && l2 != nil {
		if l1.Val > l2.Val {
			r.Next = l2
			l2 = l2.Next
		} else {
			r.Next = l1
			l1 = l1.Next
		}
		r = r.Next
	}

	if l1 != nil {
		r.Next = l1
	}

	if l2 != nil {
		r.Next = l2
	}
	return list.Next
}
