## [反转链表](<https://leetcode-cn.com/problems/reverse-linked-list/>)

###  题目

反转一个单链表。

**示例:**

```
输入: 1->2->3->4->5->NULL
输出: 5->4->3->2->1->NULL
```

**进阶：**

你可以迭代或递归地反转链表。你能否用两种方法解决这道题？

### 思路

+ 简单思路

遍历原始链表，将每个结点放到新链表的头部

参考：

​	<http://www.pianshen.com/article/517476376/>

​	<https://zxi.mytechroad.com/blog/list/leetcode-206-reverse-linked-list/>

+ 迭代方式

反转后的链表，原始链表中的每个节点，都指向其上一个节点，即：

~~~
原始链表：1->2->3->4->5->NULL
反转后：  1<-2<-3<-4<-5<-NULL
~~~

分解步骤：

~~~
0: head  1->2->3->4->5->NULL

// 暂存curr后的链表，断开curr和其后的连接，curr指向prev
1: head<-1  [2->3->4->5->NULL]
		 |	 |	[ tmp						 ]
		prev curr

// curr和prev后移
2: head<-1  [2->3->4->5->NULL]
		 		 |	|[tmp						 ]
		   prev curr
		   
// 重复上述步骤，直到curr为NULL	   
3: head<-1<-2  3->4->5->NULL
		 		 |	|
		   prev curr

4: head<-1<-2  3->4->5->NULL
		 		 		|	 |
		   		prev curr
~~~

