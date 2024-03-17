package main

import "sort"

func majorityElement(nums []int) int {
	var target = len(nums)/2 + 1
	var slow, fast int
	sort.Ints(nums)
	for fast < len(nums) {
		if nums[fast] == nums[slow] {
			fast++
		} else if nums[fast] > nums[slow] {
			if fast-slow >= target {
				return nums[slow]
			}
			slow = fast
		}
	}
	return nums[slow]
}

func _majorityElement(nums []int) int {
	var res = nums[0]
	var cnt = 1
	for i := 1; i < len(nums); i++ {
		if cnt == 0 {
			cnt = 1
			res = nums[i]
			continue
		}
		if nums[i] == res {
			cnt++
		} else {
			cnt--
		}
	}
	return res
}
