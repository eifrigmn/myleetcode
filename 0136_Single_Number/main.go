/*
题目：只出现一次的数字
	给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现两次。找出那个只出现了一次的元素。

	说明：
		你的算法应该具有线性时间复杂度。 你可以不使用额外空间来实现吗？

	示例 1:
		输入: [2,2,1]
		输出: 1
	示例 2:
		输入: [4,1,2,1,2]
		输出: 4

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/single-number
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
package _0136

import "sort"

// 思路：
//	先将数组排序，则相同的数字彼此相邻
//	由于单一元素只有一个，则以两个元素进行分组，
// 	当出现组内元素不相等时，下标小的元素为所要查找的元素
// 	若按序分组查找到最后一个元素，则其肯定为所要查找的元素
func singleNumber(nums []int) int {
	sort.Ints(nums)
	for i :=0; i<len(nums)-1;i+=2{
		if nums[i] != nums[i+1] {
			return nums[i]
		}
	}
	return nums[len(nums)-1]
}

// 参考
// 位运算^:
// 	两个相同的值取异或运算，结果为0
// 	0与任意至取异或，结果为该值本身
func singleNumber1(nums []int) int {
	ans := 0
	for _, v := range nums {
		ans ^= v
	}
	return ans
}