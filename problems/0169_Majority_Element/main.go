package _0169

import "sort"

func majorityElement(nums []int) int {
	var tmpMap = map[int]int{}
	for _, v := range nums {
		count, ok := tmpMap[v]
		if ok {
			count++
			if count > len(nums)/2 {
				return v
			} else {
				tmpMap[v] = count
			}
		} else {
			tmpMap[v] = 1
		}
	}
	return nums[0]
}

// 参考解
func majorityElement1(nums []int) int {
	var rt, count int
	for k, v := range nums {
		if k == 0 {
			rt = v
			count++
		} else {
			if v != rt {
				count--
				if count == 0 {
					rt = v
					count++
				}
			} else {
				count++
			}
		}
	}
	return rt
}

func majorityElement2(nums []int)int{
	sort.Ints(nums)
	return nums[len(nums)/2]
}