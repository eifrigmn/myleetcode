/*
题目：合并两个有序链表
	将两个有序链表合并为一个新的有序链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。
 
	示例：
		输入：1->2->4, 1->3->4
		输出：1->1->2->3->4->4
来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/merge-two-sorted-lists
*/

package src

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	//if l1 == nil {
	//	return l2
	//} else if l2 == nil {
	//	return l1
	//}
	var mergeList = new(ListNode)
	var res = mergeList
	for {
		if l1 == nil {
			res.Next = l2
			break
		} else if l2 == nil {
			res.Next = l1
			break
		} else {
			if l1.Val >= l2.Val {
				res.Next = l2
				l2 = l2.Next
			} else {
				res.Next = l1
				l1 = l1.Next
			}
			res = res.Next
		}

	}
	return mergeList.Next
}

// 最优解
func mergeTwoLists1(l1 *ListNode, l2 *ListNode) *ListNode {
	pre := &ListNode{0, nil}
	result := pre
	for l1 != nil && l2 != nil {
		if l1.Val > l2.Val {
			pre.Next = l2
			l2 = l2.Next
		} else {
			pre.Next = l1
			l1 = l1.Next
		}
		pre = pre.Next
	}
	if l1 != nil {
		pre.Next = l1
	}
	if l2 != nil {
		pre.Next = l2
	}
	return result.Next
}
