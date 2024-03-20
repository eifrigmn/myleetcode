## [合并 K 个升序链表](https://leetcode.cn/problems/merge-k-sorted-lists/description/)

### 题目

给你一个链表数组，每个链表都已经按升序排列。

请你将所有链表合并到一个升序链表中，返回合并后的链表。 

### 示例1：

~~~
输入：lists = [[1,4,5],[1,3,4],[2,6]]
输出：[1,1,2,3,4,4,5,6]
解释：链表数组如下：
[
  1->4->5,
  1->3->4,
  2->6
]
将它们合并到一个有序链表中得到。
1->1->2->3->4->4->5->6
~~~

### 示例2：

~~~
输入：lists = []
输出：[]
~~~

### 示例3：

~~~
输入：lists = [[]]
输出：[]
~~~

### 解题思路
递归方式，k个list，两两合并，经过递归调用，最终合并成一个升序单链表