package _01

func removeDuplicates(nums []int) int {
	if len(nums) < 2 {
		return len(nums)
	}
	var slow int
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[slow] {
			slow++
			nums[slow] = nums[i]
		}
	}
	return slow + 1
}
