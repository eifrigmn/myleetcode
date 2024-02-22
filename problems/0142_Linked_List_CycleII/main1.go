package _0142

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func _detectCycle(head *ListNode) *ListNode {
	var mp = make(map[*ListNode]struct{})
	for head != nil {
		if _, ok := mp[head]; ok {
			return head
		}
		mp[head] = struct{}{}
		head = head.Next
	}
	return nil
}
