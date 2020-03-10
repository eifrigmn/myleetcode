package linked_list

import "myleetcode/util"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode = util.ListNode

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	var count int
	curr := head
	for curr != nil {
		count++
		curr = curr.Next
	}
	if count == n {
		return head.Next
	}
	idx := count - n
	prv := &ListNode{0, head}
	for idx > 0 {
		prv = prv.Next
		idx--
	}
	prv.Next = prv.Next.Next
	return head
}

//============================参考============================
func removeNthFromEnd1(head *ListNode, n int) *ListNode {
	d, headIsNthFromEnd := getDaddy(head, n)

	if headIsNthFromEnd {
		// 删除head节点
		return head.Next
	}

	d.Next = d.Next.Next

	return head
}

// 获取倒数第N个节点的父节点
// 从前往后数n个节点，则后面的节点数即为从后往前数第n个节点之前的节点数
func getDaddy(head *ListNode, n int) (daddy *ListNode, headIsNthFromEnd bool) {
	daddy = head

	for head != nil {
		if n < 0 {
			daddy = daddy.Next
		}

		n--
		head = head.Next
	}

	// n==0，说明链的长度等于n
	headIsNthFromEnd = n == 0

	return
}
