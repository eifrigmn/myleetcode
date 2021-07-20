package _06

import "sort"

func missingNumber(nums []int) int {
	sort.Ints(nums)
	for k, v := range nums {
		if k != v {
			return k
		}
	}
	return len(nums)
}
