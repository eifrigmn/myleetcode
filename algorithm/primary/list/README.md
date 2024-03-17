# 单链表解题归类

## 双指针解法

### 1、合并链表

> 顺序遍历节点，比较大小后合并，注意合并前要暂存被合并的节点后的链表内容

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

~~~go
// k个升序链表两两合并，递归调用，直到合并成一个升序链表
func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	} else if len(lists) == 1 {
		return lists[0]
	}
	mid := len(lists) / 2
	return mergeTwoList(mergeKLists(lists[:mid]), mergeKLists(lists[mid:]))
}

func mergeTwoList(list1, list2 *ListNode) *ListNode {
	var result = new(ListNode)
	var head = result
	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			head.Next = list1
			list1 = list1.Next
		} else {
			head.Next = list2
			list2 = list2.Next
		}
		head = head.Next
	}
	if list1 != nil {
		head.Next = list1
	}
	if list2 != nil {
		head.Next = list2
	}
	return result.Next
}
~~~

### 2、链表的分解

分别设置两个虚拟投节点，指向分割后的两个链表，遍历原链表，把符合条件的节点分配到分割后的两个链表中

#### [86. 分隔链表](https://leetcode.cn/problems/partition-list/)

~~~go
func partition(head *ListNode, x int) *ListNode {
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
~~~

### 3、寻找单链表的倒数第 `k` 个节点

设置两个指针，p1与p2相差k个节点，当p2走到链表尾时，p1指向的节点即为链表倒数第k个节点

#### [19. 删除链表的倒数第 N 个结点](https://leetcode.cn/problems/remove-nth-node-from-end-of-list/)

~~~go
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	var q = new(ListNode)
	q.Next = head
	var p1 = q
	var p2 = q
	for i := 0; i < n+1; i++ {
		if p1 == nil {
			p1 = head
		} else {
			p1 = p1.Next
		}
	}
	for p1 != nil {
		p1 = p1.Next
		p2 = p2.Next
	}
	p2.Next = p2.Next.Next
	return q.Next
}
~~~

### 4、寻找单链表的中点

快慢指针，快指针每次移动2步，慢指针每次移动1步，快指针走完时，满指针指向的节点即为单链表的中点

#### [876. 链表的中间结点](https://leetcode.cn/problems/middle-of-the-linked-list/)

~~~go
func middleNode(head *ListNode) *ListNode {
	var slow, fast = head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}
~~~

### 5、链表中有环

快慢指针，快指针每次移动2步，慢指针每次移动1步，

+ 若二者相遇，则链表中有环
+ 二者相遇后，任意一个指针指向链表的头节点，两个指针依次每次移动1步，二者再次相遇时，则找到环的起点

[141. 环形链表](https://leetcode.cn/problems/linked-list-cycle/)

~~~go
func hasCycle(head *ListNode) bool {
	var slow, fast = head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return true
		}
	}
	return false
}
~~~

#### [142. 环形链表 II](https://leetcode.cn/problems/linked-list-cycle-ii/)

~~~go
func detectCycle(head *ListNode) *ListNode {
	var crossNode = getCrossNode(head)
	if crossNode == nil {
		return nil
	}
	var h = head
	for h != nil && crossNode != nil {
		if h == crossNode {
			return crossNode
		}
		h = h.Next
		crossNode = crossNode.Next
	}
	return nil
}

func getCrossNode(head *ListNode) *ListNode {
	var slow, fast = head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return fast
		}
	}
	return nil
}
~~~

### 6、链表是否相交

设置两个指针p1和p2，分别指向两个链表，依次遍历，p1或者p2遍历完当前链表时，将其指向另外一条链表继续遍历，当p1=p2时，则该节点为链表的相交节点。

#### [160. 相交链表](https://leetcode.cn/problems/intersection-of-two-linked-lists/)

~~~go
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	var p1 = headA
	var p2 = headB
	for p1 != p2 {
		if p1 == nil {
			p1 = headB
		} else {
			p1 = p1.Next
		}
		if p2 == nil {
			p2 = headA
		} else {
			p2 = p2.Next
		}
	}
	return p1
}
~~~

## 单链表实现LRU算法

