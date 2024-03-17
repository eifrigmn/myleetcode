package main

func merge(nums1 []int, m int, nums2 []int, n int) {
	if n == 0 {
		return
	}
	if m == 0 {
		copy(nums1, nums2)
		return
	}
	var i, j int
	for i < m+j && j < n {
		if nums1[i] > nums2[j] {
			moveOneStep(nums1, i)
			nums1[i] = nums2[j]
			j++
		}
		i++
	}
	for j < n {
		nums1[i] = nums2[j]
		j++
		i++
	}
}

// 从指定位置向右移动元素
func moveOneStep(nums []int, idx int) {
	for i := len(nums) - 2; i >= idx; i-- {
		nums[i+1] = nums[i]
	}
}
