package util

import (
	"fmt"
	"myleetcode/linked_list"
)

func ArryToSinglyList(nums []int) *linked_list.ListNode {
	head := &linked_list.ListNode{}
	result := head
	for i := 0; i < len(nums); i++ {
		head.Next = &linked_list.ListNode{Val: nums[i]}
		head = head.Next
	}
	return result.Next
}

func FmtList(list interface{}) {
	l := list.(*linked_list.ListNode)
	var output string
	if l == nil {
		fmt.Println("空链表")
		return
	}
	for l.Next != nil {
		output = fmt.Sprintf("%s, %d", output, l.Val)
		l = l.Next
	}
	fmt.Println(output)
}
