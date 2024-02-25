package _0021

func _mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	// 设置一个虚拟头结点
	var result = new(ListNode)
	var head = result
	// 遍历list1和list2，其中任意一个遍历完毕，则跳出循环
	for list1 != nil && list2 != nil {
		// 拿出较小的结点拼接到结果中
		if list1.Val < list2.Val {
			head.Next = list1
			list1 = list1.Next
		} else {
			head.Next = list2
			list2 = list2.Next
		}
		head = head.Next
	}
	// 可能存在某条链表未遍历完毕，则直接将剩余的子链表直接拼接到结果
	if list1 != nil {
		head.Next = list1
	} else if list2 != nil {
		head.Next = list2
	}
	// 注意要去除虚拟头节点
	return result.Next
}
