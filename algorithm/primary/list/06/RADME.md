# [环形链表](https://leetcode-cn.com/explore/interview/card/top-interview-questions-easy/6/linked-list/46/)

## 题目

给定一个链表，判断链表中是否有环。

为了表示给定链表中的环，我们使用整数 pos 来表示链表尾连接到链表中的位置（索引从 0 开始）。 如果 pos 是 -1，则在该链表中没有环。

### 示例

+ 示例1

```
输入：head = [3,2,0,-4], pos = 1
输出：true
解释：链表中有一个环，其尾部连接到第二个节点。
```

![](../../../../assets/circularlinkedlist_test1.png)

+ 示例2

```
输入：head = [1,2], pos = 0
输出：true
解释：链表中有一个环，其尾部连接到第一个节点。
```

![](../../../../assets/circularlinkedlist_test2.png)

+ 示例3

```
输入：head = [1], pos = -1
输出：false
解释：链表中没有环。
```

![](../../../../assets/circularlinkedlist_test3.png)

### 进阶

你能用`O(1)`(即，常量)内存解决此问题吗？

### 思路

参考如下图

![](../../../../assets/142_solution.png)