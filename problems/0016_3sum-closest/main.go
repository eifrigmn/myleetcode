package _016

import (
	"sort"
)

func threeSumClosest(nums []int, target int) int {
	var step = 10000
	var result int
	sort.Ints(nums)
	for i:=0;i<len(nums);i++{
		lft := i+1
		rgt := len(nums)-1
		for lft < rgt {
			sum := nums[i]+nums[lft]+nums[rgt]
			if sum == target {
				return target
			} else if sum < target {
				if step > target-sum {
					step = target-sum
					result = sum
				}
				lft++
			} else {
				if step > sum-target {
					step = sum-target
					result = sum
				}
				rgt--
			}
		}
	}
	return result
}
