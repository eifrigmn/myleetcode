package _01

func merge(nums1 []int, m int, nums2 []int, n int) {
	for i := 0; i < m && len(nums2) > 0; i++ {
		if nums1[i] >= nums2[0] {
			copy(nums1[i+1:m+1], nums1[i:m])
			nums1[i] = nums2[0]
			nums2 = nums2[1:]
			m++
			n--
		}
	}
	if n > 0 {
		copy(nums1[m:m+n], nums2[:])
	}
}
