package _0083

func _partition(head *ListNode, x int) *ListNode {
	var result = new(ListNode)
	var lft = result
	var rgt = new(ListNode)
	var r = rgt
	for head != nil {
		tmp := head.Next
		if head.Val < x {
			lft.Next = head
			lft = lft.Next
		} else {
			r.Next = head
			r = r.Next
			r.Next = nil
		}
		head = tmp
	}
	lft.Next = rgt.Next
	return result.Next
}
