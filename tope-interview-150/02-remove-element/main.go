package main

func removeElement(nums []int, val int) int {
	if len(nums) == 1 && nums[0] == val {
		nums = []int{}
		return 0
	}
	var i = 0
	var size = len(nums)
	for i < size {
		if nums[i] == val {
			copy(nums[i:size-1], nums[i+1:size])
			size--
		} else {
			i++
		}
	}
	nums = nums[:size]
	return size
}

func _removeElement(nums []int, val int) int {
	p := -1
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			p++
			nums[p] = nums[i]
		}
	}
	return p + 1
}

func removeElement1(nums []int, val int) int {
	var slow, fast int
	for fast < len(nums) {
		if nums[fast] != val {
			nums[slow] = nums[fast]
			slow++
		}
		fast++
	}
	return slow
}
