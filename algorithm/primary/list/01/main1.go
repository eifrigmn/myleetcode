package _01

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteNode1(node *ListNode) {
	node.Val = node.Next.Val
	node.Next = node.Next.Next
}
