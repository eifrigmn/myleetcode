package _0206

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func _reverseList(head *ListNode) *ListNode {
	var result = new(ListNode)
	for head != nil {
		tmp := result.Next
		nodeVal := head.Val
		result.Next = &ListNode{
			Val:  nodeVal,
			Next: tmp,
		}
		head = head.Next
	}
	return result.Next
}
