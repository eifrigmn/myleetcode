## [合并区间](<https://leetcode-cn.com/problems/merge-intervals/>)

### 题目

给出一个区间的集合，请合并所有重叠的区间。

**示例1：**

~~~
输入: [[1,3],[2,6],[8,10],[15,18]]
输出: [[1,6],[8,10],[15,18]]
解释: 区间 [1,3] 和 [2,6] 重叠, 将它们合并为 [1,6].
~~~

**示例2：**

~~~
输入: [[1,4],[4,5]]
输出: [[1,5]]
解释: 区间 [1,4] 和 [4,5] 可被视为重叠区间。
~~~

### 思路

先将二维数组按照第一个值的大小进行升序排序，第一个值相等的，按照第二个值升序排序(参考快排思路)

遍历二维数组，顺序比较下一个元素的第一个值是否在上一个元素的区间内，如果在，则合并两个元素的区间

