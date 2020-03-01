package util

type ListNode struct {
	Val  int
	Next *ListNode
}

// SListToInts 单链表转换为数组
func SListToInts(head *ListNode) []int {
	var nums []int
	for head != nil {
		nums = append(nums, head.Val)
		head = head.Next
	}
	return nums
}

// IntsToSList 数组转化为单链表
func IntsToSList(nums []int) *ListNode {
	head := &ListNode{}
	result := head
	for i := 0; i < len(nums); i++ {
		head.Next = &ListNode{Val: nums[i]}
		head = head.Next
	}
	return result.Next
}

// Ints2ListWithCycle returns a list whose tail point to pos-indexed node
// head's index is 0
// if pos = -1, no cycle
func Ints2ListWithCycle(nums []int, pos int) *ListNode {
	head := IntsToSList(nums)
	if pos == -1 {
		return head
	}
	c := head
	for pos > 0 {
		c = c.Next
		pos--
	}
	tail := c
	for tail.Next != nil {
		tail = tail.Next
	}
	tail.Next = c
	return head
}

// head must Not be nil
func TailOf(head *ListNode) *ListNode {
	for head.Next != nil {
		head = head.Next
	}
	return head
}
