package main

func removeDuplicates(nums []int) int {
	var slow, fast int
	for fast < len(nums) {
		if nums[fast] > nums[slow] {
			slow++
			nums[slow] = nums[fast]
		}
		fast++
	}
	return slow + 1
}
