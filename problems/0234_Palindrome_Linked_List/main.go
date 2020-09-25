package _0234

import "myleetcode/util"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode = util.ListNode

func isPalindrome1(head *ListNode) bool {
	arr := []int{}
	for head != nil {
		arr = append(arr, head.Val)
		head = head.Next
	}

	for i := 0; i < len(arr)/2; i++ {
		if arr[i] != arr[len(arr)-1-i] {
			return false
		}
	}

	return true
}

// 翻转链表然后顺序遍历
func isPalindrome(head *ListNode) bool {
	n := 0
	for p := head; p != nil; p = p.Next {
		n++
	}
	// 找到链表中点的下一个位置
	p := head
	for i := 0; i < n/2; i++ {
		p = p.Next
	}
	if n%2 == 1 {
		p = p.Next
	}
	// 翻转后半段链表
	p = reverse(p)
	for p != nil {
		if head.Val != p.Val {
			return false
		}
		p = p.Next
		head = head.Next
	}
	return true
}

// 反转一个链表，并返回反转后到头节点
func reverse(head *ListNode) *ListNode {
	var prev *ListNode
	for head != nil {
		head.Next, prev, head = prev, head, head.Next
	}
	return prev
}
