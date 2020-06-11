/*
题目：缺失的第一个正数
	给定一个未排序的整数数组，找出其中没有出现的最小的正整数。

	示例 1:
		输入: [1,2,0]
		输出: 3
	示例 2:
		输入: [3,4,-1,1]
		输出: 2
	示例 3:
		输入: [7,8,9,11,12]
		输出: 1

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/first-missing-positive
*/
package _0041

import (
	"sort"
)

func firstMissingPositive(nums []int) int {
	sort.Ints(nums)
	for i := 0; i < len(nums); i++ {
		if i == len(nums)-1 {
			if nums[i] == i+1 {
				return len(nums) + 1
			}
			return i + 1
		}

		if nums[i] <= 0 {
			nums = nums[i+1:]
			i--
			continue
		}

		if nums[i] == nums[i+1] {
			nums = append(nums[:i], nums[i+1:]...)
			i--
			continue
		}

		if nums[i] != i+1 {
			return i + 1
		}
	}
	return len(nums) + 1
}

// 参考解
func firstMissingPositive1(nums []int) int {
	haveOne := false
	for i, n := range nums {
		if n == 1 {
			haveOne = true
		}
		if n <= 0 || n > len(nums) {
			nums[i] = 1
		}
	}
	if !haveOne {
		return 1
	}
	if len(nums) == 1 {
		return 2
	}
	for _, n := range nums {
		if n < 0 {
			n = -n
		}
		n = n - 1
		if nums[n] > 0 {
			nums[n] = -nums[n]
		}
	}
	ret := 0
	for ret < len(nums) {
		if nums[ret] > 0 {
			return ret + 1
		}
		ret++
	}
	return ret + 1
}