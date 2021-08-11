package _002

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// 设立一个全局变量，标记是否有进位
// 循环退出的标记为：遍历完两个链表l1、l2以及进位标志不为1
func addTwoNumbers1(l1 *ListNode, l2 *ListNode) *ListNode {
	var head = &ListNode{}
	var add int
	var h = head
	for l1 != nil || l2 != nil || add != 0 {
		num := add
		if l1 != nil {
			num += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			num += l2.Val
			l2 = l2.Next
		}
		add = num / 10
		num = num % 10
		h.Next = &ListNode{Val: num}
		h = h.Next
	}
	return head.Next
}
