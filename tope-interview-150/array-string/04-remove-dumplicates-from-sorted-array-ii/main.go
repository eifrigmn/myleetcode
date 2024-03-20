package main

func removeDuplicates(nums []int) int {
	if len(nums) < 3 {
		return len(nums)
	}

	var slow, fast = 0, 1
	for fast < len(nums) {
		if nums[slow] == nums[fast] {
			if slow >= 1 && nums[slow] == nums[slow-1] {
				fast++
				continue
			}
		}
		slow++
		nums[slow] = nums[fast]
		fast++
	}
	return slow + 1
}

func _removeDuplicates(nums []int) int {
	n := len(nums)
	if n <= 2 {
		return n
	}

	slow, fast := 2, 2
	for fast < n {
		if nums[slow-2] != nums[fast] {
			nums[slow] = nums[fast]
			slow++
		}
		fast++
	}
	return slow
}
