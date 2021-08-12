package _0004

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestFindMedianSortedArrays(t *testing.T) {
	Convey("findMedianSortedArrays", t, func() {
		nums1 := []int{1, 3}
		nums2 := []int{2}
		target := 2.00000
		result := findMedianSortedArrays1(nums1, nums2)
		So(result, ShouldEqual, target)

		nums1 = []int{1, 2}
		nums2 = []int{3, 4}
		target = 2.50000
		result = findMedianSortedArrays1(nums1, nums2)
		So(result, ShouldEqual, target)

		nums1 = []int{0, 0}
		nums2 = []int{0, 0}
		target = 0.00000
		result = findMedianSortedArrays1(nums1, nums2)
		So(result, ShouldEqual, target)

		nums1 = []int{}
		nums2 = []int{1}
		target = 1.00000
		result = findMedianSortedArrays1(nums1, nums2)
		So(result, ShouldEqual, target)

		nums1 = []int{2}
		nums2 = []int{}
		target = 2.00000
		result = findMedianSortedArrays1(nums1, nums2)
		So(result, ShouldEqual, target)
	})
}
