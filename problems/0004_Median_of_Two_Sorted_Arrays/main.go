package _0004

import (
	"fmt"
	"strconv"
)

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	m := len(nums1)
	n := len(nums2)
	var isEven bool
	if (m+n)%2 == 0 {
		isEven = true
	}

	var mid = (m+n)/2
	var target int
	for i:=0;i<mid-1;i++{
		_, nums1, nums2 = giveNext(nums1, nums2)
	}

	var n1 int
	n1, nums1, nums2 = giveNext(nums1, nums2)
	target, _, _ = giveNext(nums1, nums2)
	if isEven || m+n <= 2 {
		target += n1
	}

	value := float64(target)
	if isEven {
		value = value / 2
	}
	value, _ = strconv.ParseFloat(fmt.Sprintf("%.5f", value), 64)
	return value
}

func giveNext(nums1, nums2 []int) (int, []int, []int){
	var result int
	if len(nums1) == 0 && len(nums2) == 0 {
		return 0, nil, nil
	} else if len(nums1) == 0 {
		result = nums2[0]
		nums2 = nums2[1:]
	} else if len(nums2) == 0 {
		result = nums1[0]
		nums1 = nums1[1:]
	} else {
		if nums1[0] > nums2[0] {
			result = nums2[0]
			nums2 = nums2[1:]
		} else {
			result = nums1[0]
			nums1 = nums1[1:]
		}
	}
	return result, nums1, nums2
}