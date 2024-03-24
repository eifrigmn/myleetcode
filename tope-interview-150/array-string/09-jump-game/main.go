package main

func canJump(nums []int) bool {
	if len(nums) < 2 {
		return true
	}
	// 以倒数第二位开始倒序遍历，倒数第二位跳到最后一位，需要跳跃的最小步数为1
	var minStep = 1
	for i := len(nums) - 2; i >= 0; i-- {
		// 第i位可以跳跃的最大长度小于minStep时，则不是从第i位跳跃而来的，则minStep+1
		if nums[i] < minStep {
			minStep++
			continue
		}
		minStep = 1
	}
	// 最终最小步数应该为1
	return minStep == 1
}
