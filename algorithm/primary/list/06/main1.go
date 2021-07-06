package _06

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func hasCycle1(head *ListNode) bool {
	var fast = head
	var slow = head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return true
		}
	}
	return false
}

func hasCycle2(head *ListNode) bool {
	if head == nil {
		return false
	}
	for {
		tmp := head.Next
		head.Next = head
		head = tmp
		if head == nil {
			return false
		}
		if head == head.Next {
			return true
		}
	}
}
