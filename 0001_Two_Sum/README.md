# [两数之和](https://leetcode-cn.com/problems/two-sum)
## 题目	

给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那两个整数，并返回他们的数组下标。

​	你可以假设每种输入只会对应一个答案。但是，你不能重复利用这个数组中同样的元素。

### 示例:

	给定 nums = [2, 7, 11, 15], target = 9
	因为 nums[0] + nums[1] = 2 + 7 = 9
	所以返回 [0, 1]
### 视频题解：

[花花酱 LeetCode 1. Two Sum - 刷题找工作 EP1](https://www.youtube.com/watch?v=tNtk_rwbaIk&list=PLLuMmzMTgVK7eabhqdNs1wQv0cFEbd048&index=58)

### 解题思路：

如果数组中存在两个值`a`和`b`相加之和为`target`，则：

~~~
a + b = targer
~~~

即：

~~~
b = target - a
~~~

设置一个map存储`nums`数组中的值，其中：key为nums中的元素，value为对应元素的游标

遍历`nums`，如果从`map`中找到对应的`target - nums[i]`，则返回对应两个值的游标，否则，将该元素存放于`map`中。