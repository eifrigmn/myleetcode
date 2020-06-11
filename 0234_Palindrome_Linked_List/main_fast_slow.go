package _0234

func isPalindrome2(head *ListNode) bool {
	var nums []int
	fast, slow := head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		nums = append(nums, slow.Val)
		slow = slow.Next
	}
	if fast != nil && fast.Next == nil {
		// 节点为奇数个
		slow = slow.Next
	}
	for slow != nil && slow.Val == nums[len(nums)-1] {
		slow = slow.Next
		nums = nums[:len(nums)-1]
	}
	return len(nums) == 0
}
