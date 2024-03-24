package main

func jump(nums []int) int {
	if len(nums) == 1 {
		return 0
	}
	if nums[0] >= len(nums)-1 {
		return 1
	}
	var result int
	for i := 0; i < len(nums)-1; {
		nextI := getMaxStartIdx(nums, i)
		i = nextI
		result++
	}
	return result
}

// 以i为起点，寻找其覆盖范围内的下一个能跳到最远距离的起点
func getMaxStartIdx(nums []int, i int) int {
	maxStep := nums[i]
	if maxStep >= len(nums)-i-1 {
		return len(nums) - 1
	}
	var maxIdx = i + 1
	for j := i + 1; j <= i+maxStep && j < len(nums); j++ {
		if nums[j]+j > nums[maxIdx]+maxIdx {
			maxIdx = j
		}
	}
	return maxIdx
}
