package _0088

import "testing"

func TestMerge(t *testing.T) {
	var (
		nums1 = []int{1, 2, 3, 0, 0, 0, 0, 0}
		nums2 = []int{2, 5, 6}
		m     = 3
		n     = 3
	)
	merge(nums1, m, nums2, n)
}
