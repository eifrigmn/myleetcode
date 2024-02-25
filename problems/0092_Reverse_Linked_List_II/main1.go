package _0206

func _reverseBetween(head *ListNode, left int, right int) *ListNode {
	var i = 1
	var prev *ListNode
	var curr = head
	for curr != nil {
		if i > left && i <= right {
			tmp := curr.Next
			curr.Next = prev
			prev = curr
			curr = tmp
		} else {
			prev = curr
			curr = curr.Next
		}
		i++
	}
	return head
}
