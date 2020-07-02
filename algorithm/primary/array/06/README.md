# [两个数组的交集 II](https://leetcode-cn.com/explore/interview/card/top-interview-questions-easy/1/array/26/)

## 题目

给定两个数组，编写一个函数来计算它们的交集。

### 示例

+ 示例1：

```
输入: nums1 = [1,2,2,1], nums2 = [2,2]
输出: [2,2]
```

+ 示例2：

```
输入: nums1 = [4,9,5], nums2 = [9,4,9,8,4]
输出: [4,9]
```

### 说明

+ 输出结果中每个元素出现的次数，应与元素在两个数组中出现的次数一致。
+ 我们可以不考虑输出结果的顺序。

### 进阶

+ 如果给定的数组已经排好序呢？你将如何优化你的算法？
+ 如果 nums1 的大小比 nums2 小很多，哪种方法更优？
+ 如果 nums2 的元素存储在磁盘上，磁盘内存是有限的，并且你不能一次加载所有的元素到内存中，你该怎么办？

### 思路

#### 思路一

遍历**nums1**并在**nums2**中寻找相同值，如果找到，则将该值存储于结果集中，并将该值从**nums2**中删除。

#### 思路二

将**nums1**中的元素存储于map中，其中key为元素的值，value为元素出现的次数
遍历**nums2**如果在map中找到相应的值，则对应值的value减一，并把该值存储于结果集中
