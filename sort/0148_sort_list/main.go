package sort

import "myleetcode/util"

type ListNode = util.ListNode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	// 链表中点
	lft, rgt := split(head)
	return mergeSort(sortList(lft), sortList(rgt))
}

func split(head *ListNode) (lft, rgt *ListNode) {
	var slow, fast = head, head
	var slowPrv *ListNode
	for fast != nil && fast.Next != nil {
		slowPrv, slow = slow, slow.Next
		fast = fast.Next.Next
	}
	slowPrv.Next = nil
	lft, rgt = head, slow
	return
}

func mergeSort(lft, rgt *ListNode) *ListNode {
	h := &ListNode{}
	result := h
	for lft != nil && rgt != nil {
		if lft.Val <= rgt.Val {
			h.Next, lft = lft, lft.Next
		} else {
			h.Next, rgt = rgt, rgt.Next
		}
		h = h.Next
	}

	if lft == nil {
		h.Next = rgt
	}

	if rgt == nil {
		h.Next = lft
	}
	return result.Next
}
