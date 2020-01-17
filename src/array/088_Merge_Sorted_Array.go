package array

import (
	"fmt"
	"sort"
)

func merge(nums1 []int, m int, nums2 []int, n int) {
	nums := append(nums1[:m], nums2[:n]...)
	sort.Ints(nums)
	copy(nums1[:n+m], nums)
	fmt.Println(nums1)
}
