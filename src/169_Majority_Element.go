/*
题目：多数元素
	给定一个大小为 n 的数组，找到其中的多数元素。多数元素是指在数组中出现次数大于 ⌊ n/2 ⌋ 的元素。
	你可以假设数组是非空的，并且给定的数组总是存在多数元素。

	示例 1:
		输入: [3,2,3]
		输出: 3
	示例 2:
		输入: [2,2,1,1,1,2,2]
		输出: 2

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/majority-element
*/
package src

import "sort"

func majorityElement(nums []int) int {
	var tmpMap = map[int]int{}
	for _, v := range nums {
		count, ok := tmpMap[v]
		if ok {
			count++
			if count > len(nums)/2 {
				return v
			} else {
				tmpMap[v] = count
			}
		} else {
			tmpMap[v] = 1
		}
	}
	return nums[0]
}

// 参考解
func majorityElement1(nums []int) int {
	var rt, count int
	for k, v := range nums {
		if k == 0 {
			rt = v
			count++
		} else {
			if v != rt {
				count--
				if count == 0 {
					rt = v
					count++
				}
			} else {
				count++
			}
		}
	}
	return rt
}

func majorityElement2(nums []int)int{
	sort.Ints(nums)
	return nums[len(nums)/2]
}