package linked_list

import (
	"fmt"
	"testing"
)

func TestDeleteDuplicates(t *testing.T) {
	// 1->1->2
	l1 := arryToSinglyList([]int{1,1,2})
	newl1 := deleteDuplicates(l1)
	fmtList(newl1)
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

	for l != nil {
		fmt.Printf(" %d ", l.Val)
		l = l.Next
	}
}
