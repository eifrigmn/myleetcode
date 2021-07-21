package _01

func removeDuplicates1(nums []int) int {
	if len(nums) < 2 {
		return len(nums)
	}
	var l = 1
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[i-1] {
			nums[l] = nums[i]
			l++
		}
	}
	return l
}
