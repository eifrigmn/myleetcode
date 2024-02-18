package _0001

func _twoSum(nums []int, target int) []int {
	var tmp = make(map[int]int, len(nums))
	for i := 0; i < len(nums); i++ {
		if j, ok := tmp[target-nums[i]]; ok {
			return []int{i, j}
		}
		tmp[nums[i]] = i
	}
	return []int{-1, -1}
}
