/*
题目：三数之和

	给定一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？找出所有满足条件且不重复的三元组。
	注意：答案中不可以包含重复的三元组。

	例如, 给定数组 nums = [-1, 0, 1, 2, -1, -4]，
	满足要求的三元组集合为：
	[
  		[-1, 0, 1],
  		[-1, -1, 2]
	]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/3sum
*/
package array

import "sort"

/*
参考解题思路：
	特判，对于数组长度 n，如果数组为 nil 或者数组长度小于 3，返回 nil。
	对数组进行排序。
	遍历排序后数组：
		若 nums[i]>0：因为已经排序好，所以后面不可能有三个数加和等于 0，直接返回结果。
		对于重复元素：跳过，避免出现重复解
		令左指针 L=i+1，右指针 R=n-1，当 L<R 时，执行循环：
			当 nums[i]+nums[L]+nums[R]==0，执行循环，判断左界和右界是否和下一位置重复，去除重复解。并同时将 L,R 移到下一位置，寻找新的解
			若和大于 0，说明 nums[R] 太大，R 左移
			若和小于 0，说明 nums[L] 太小，L 右移
	复杂度分析
		时间复杂度：O(n2), 数组排序 O(nlogn)，遍历数组 O(n)，双指针遍历 O(n)，总体 O(nlogn) + O(n) * O(n), O(n2)
		空间复杂度：O(1)

作者：zhu_shi_fu
链接：https://leetcode-cn.com/problems/3sum/solution/pai-xu-shuang-zhi-zhen-zhu-xing-jie-shi-python3-by/
来源：力扣（LeetCode）
*/
func threeSum(nums []int) [][]int{
	if len(nums) <3{
		return nil
	}
	// 排序
	sort.Ints(nums)
	var res [][]int
	for i:=0;i<len(nums);i++{
		if nums[i] >0 {
			return res
		}
		if i>0 && (nums[i] == nums[i-1]) {
			continue
		}
		lftIndex := i+1
		rgtIndex := len(nums)-1
		for lftIndex<rgtIndex {
			if nums[i]+nums[lftIndex]+nums[rgtIndex] == 0 {
				res = append(res, []int{nums[i], nums[lftIndex], nums[rgtIndex]})
				lftIndex++
				rgtIndex--
				for lftIndex<rgtIndex && nums[lftIndex] == nums[lftIndex-1] {
					lftIndex++
				}

				for lftIndex<rgtIndex && nums[rgtIndex] == nums[rgtIndex+1] {
					rgtIndex--
				}
			} else if nums[i]+nums[lftIndex]+nums[rgtIndex] >0 {
				rgtIndex--
			} else {
				lftIndex++
			}
		}
	}
	return res
}
