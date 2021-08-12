package _0004

// 两个切片看作是一个合并后的升序切片
// 先按顺序丢弃掉中位点之前的元素
// 根据切片原始长度奇偶性，中位点可能是一个或者两个
// 取出中位点元素，如果是两个元素，求平均值
func findMedianSortedArrays1(nums1 []int, nums2 []int) float64 {
	var size = (len(nums1) + len(nums2))
	for len(nums1)+len(nums2) > size/2+1 {
		_, nums1, nums2 = findMin(nums1, nums2)
	}

	var tmp int
	tmp, nums1, nums2 = findMin(nums1, nums2)
	var result = float64(tmp)
	if size%2 == 0 {
		tmp, _, _ = findMin(nums1, nums2)
		result = (result + float64(tmp)) / 2
	}
	return result
}

func findMin(nums1, nums2 []int) (int, []int, []int) {
	var result int
	if len(nums1) > 0 && len(nums2) > 0 {
		if nums1[0] > nums2[0] {
			result = nums2[0]
			nums2 = nums2[1:]
		} else {
			result = nums1[0]
			nums1 = nums1[1:]
		}
	} else if len(nums1) > 0 {
		result = nums1[0]
		nums1 = nums1[1:]
	} else if len(nums2) > 0 {
		result = nums2[0]
		nums2 = nums2[1:]
	}
	return result, nums1, nums2
}
