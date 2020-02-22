package linked_list

import (
	"fmt"
	"testing"
)

func TestMergeTwoLists(t *testing.T) {
	// 1->2->4
	nums1 := []int{1, 2, 4}
	l1 := arryToSinglyList(nums1)
	// 1->3->4
	nums2 := []int{1, 3, 4}
	l2 := arryToSinglyList(nums2)
	result := mergeTwoLists(l1, l2)
	fmtList(result)
}

func arryToSinglyList(nums []int) *ListNode {
	head := &ListNode{}
	result := head
	for i := 0; i < len(nums); i++ {
		head.Next = &ListNode{Val: nums[i]}
		head = head.Next
	}
	return result.Next
}

func fmtList(l *ListNode) {
	if l == nil {
		fmt.Println("空链表")
		return
	}

	for l.Next != nil {
		fmt.Printf(" %d ", l.Val)
		l = l.Next
	}
}
