package _05

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func isPalindrome1(head *ListNode) bool {
	l, rl := reverse(head)
	for l != nil {
		if l.Val != rl.Val {
			return false
		}
		l = l.Next
		rl = rl.Next
	}
	return true
}

func reverse(head *ListNode) (*ListNode, *ListNode) {
	var l = new(ListNode)
	lh := l
	var rl = new(ListNode)
	for head != nil {
		lh.Next = new(ListNode)
		lh.Next.Val = head.Val
		lh = lh.Next
		tmp := head.Next
		head.Next = rl.Next
		rl.Next = head
		head = tmp
	}
	return l.Next, rl.Next
}
