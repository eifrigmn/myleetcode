package _0026

func _removeDuplicates(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}
	var slow, fast = 0, 1
	for fast < len(nums) {
		if nums[fast] != nums[slow] {
			slow++
			nums[slow] = nums[fast]
		}
		fast++
	}
	nums = nums[:slow+1]
	return slow + 1
}
