# 单链表解题归类

## 双指针解法

### 1、合并两个有序链表

#### [21. 合并两个有序链表](https://leetcode.cn/problems/merge-two-sorted-lists/)

~~~go
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
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
~~~

#### [23. 合并K个升序链表](https://leetcode.cn/problems/merge-k-sorted-lists/)



2、链表的分解

3、合并 `k` 个有序链表

4、寻找单链表的倒数第 `k` 个节点

5、寻找单链表的中点

6、判断单链表是否包含环并找出环起点

7、判断两个单链表是否相交并找出交点

| leetcode题                                                   | 知识点   | 难度 |
| ------------------------------------------------------------ | -------- | ---- |
| [141. 环形链表](https://leetcode.cn/problems/linked-list-cycle/) | 快慢指针 | 简单 |
| [142. 环形链表 II](https://leetcode.cn/problems/linked-list-cycle-ii/) | 快慢指针 | 中等 |
| [160. 相交链表](https://leetcode.cn/problems/intersection-of-two-linked-lists/) |          |      |
| [19. 删除链表的倒数第 N 个结点](https://leetcode.cn/problems/remove-nth-node-from-end-of-list/) |          |      |
| [21. 合并两个有序链表](https://leetcode.cn/problems/merge-two-sorted-lists/) |          |      |
| [23. 合并K个升序链表](https://leetcode.cn/problems/merge-k-sorted-lists/) |          |      |
| [86. 分隔链表](https://leetcode.cn/problems/partition-list/) |          |      |
| [876. 链表的中间结点](https://leetcode.cn/problems/middle-of-the-linked-list/) |          |      |



## 递归解法

## 如何K个一组反转链表

## 如何判断回文链表





链表需要注意相关的指针操作，即链表头和链表节点的区别

## 快慢指针法
如查找链表的中点
判断链表中是否有环
## 双链表
如翻转链表
## 单链表实现LRU算法

