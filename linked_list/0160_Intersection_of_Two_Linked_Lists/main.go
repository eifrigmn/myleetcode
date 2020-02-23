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

func getIntersectionNode1(headA, headB *ListNode) *ListNode {
	var la, lb int
	a, b := headA, headB
	// 分别求两个链表的长度
	for a != nil || b != nil {
		if a != nil {
			la++
			a = a.Next
		}
		if b != nil {
			lb++
			b = b.Next
		}
	}
	a, b = headA, headB
	if la > lb {
		for i := 0; i < la-lb; i++ {
			a = a.Next
		}
	} else if la < lb {
		for i := 0; i < lb-la; i++ {
			b = b.Next
		}
	}
	for a != nil {
		if a == b {
			return a
		}
		a = a.Next
		b = b.Next
	}
	return nil
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	a, b := headA, headB
	for a != nil || b != nil {
		if a == b {
			return a
		}
		if a == nil {
			a = headB
		}

		if b == nil {
			b = headA
		}
		a = a.Next
		b = b.Next
	}
	return nil
}

func getIntersectionNode3(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	pa := headA
	pb := headB
	for headA != nil || headB != nil {
		if headA == nil {
			headA = pb
		}

		if headB == nil {
			headB = pa
		}

		if headA == headB {
			return headA
		}

		headA = headA.Next
		headB = headB.Next
	}

	return nil
}
