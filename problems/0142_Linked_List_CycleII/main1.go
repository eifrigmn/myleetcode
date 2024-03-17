package _0142

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func _detectCycle(head *ListNode) *ListNode {
	var crossNode = getCrossNode(head)
	if crossNode == nil {
		return nil
	}
	var h = head
	for h != nil && crossNode != nil {
		if h == crossNode {
			return crossNode
		}
		h = h.Next
		crossNode = crossNode.Next
	}
	return nil
}

func getCrossNode(head *ListNode) *ListNode {
	var slow, fast = head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return fast
		}
	}
	return nil
}
