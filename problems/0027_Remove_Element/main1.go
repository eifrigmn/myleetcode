package _0027

func _removeElement(nums []int, val int) int {
	if len(nums) == 0 {
		return 0
	}
	var slow, fast int
	for fast < len(nums) {
		if nums[fast] != val {
			nums[slow] = nums[fast]
			slow++
		}
		fast++
	}
	nums = nums[:slow]
	return slow
}
