package _05

import "myleetcode/util"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode = util.ListNode
func isPalindrome(head *ListNode) bool {
	var fast, slow = head, head
	var nums []int
	for fast != nil && fast.Next != nil {
		nums = append(nums, slow.Val)
		slow = slow.Next
		fast = fast.Next.Next
	}
	// 链表长度为奇数
	if fast != nil && fast.Next == nil {
		slow = slow.Next
	}

	for slow != nil && slow.Val == nums[len(nums)-1] {
		slow = slow.Next
		nums = nums[:len(nums)-1]
	}
	return len(nums) == 0
}
