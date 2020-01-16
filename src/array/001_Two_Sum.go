/*
题目：两数之和
	给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那两个整数，并返回他们的数组下标。
	你可以假设每种输入只会对应一个答案。但是，你不能重复利用这个数组中同样的元素。

	示例:
		给定 nums = [2, 7, 11, 15], target = 9
		因为 nums[0] + nums[1] = 2 + 7 = 9
		所以返回 [0, 1]
来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/two-sum
*/
package array

func twoSum(nums []int, target int) []int {
	for i := range nums {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return []int{-1, -1}
}

// 最优解
func twoSum1(nums []int, target int) []int {
	tmpMap := make(map[int]int, len(nums))
	for i := 0; i < len(nums); i++ {
		n := target - nums[i]
		if val, ok := tmpMap[n]; ok {
			return []int{val, i}
		}
		tmpMap[nums[i]] = i
	}
	return []int{-1, -1}
}
