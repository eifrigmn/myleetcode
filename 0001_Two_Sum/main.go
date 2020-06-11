package _0001

// 双重循环，时间复杂度为O(n^2)，空间复杂度O(1)
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

// 改进后
// 使用map，key为目标值，value为目标值对应的下标
// 遍历完毕数组还未找到目标值，则返回[]int{-1,-1}
// 以空间换时间：时间复杂度为O(n)、空间复杂度为O(n)
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
