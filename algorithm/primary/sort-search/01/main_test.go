package _01

import (
	"reflect"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestMerge(t *testing.T) {
	Convey("合并有序数组", t, func() {
		// nums1 = [1,2,3,0,0,0], m = 3, nums2 = [2,5,6], n = 3
		// []int{1,2,2,3,5,6}
		nums1 := []int{1, 2, 3, 0, 0, 0}
		m := 3
		nums2 := []int{2, 5, 6}
		n := 3
		merge(nums1, m, nums2, n)
		So(reflect.DeepEqual(nums1, []int{1, 2, 2, 3, 5, 6}), ShouldBeTrue)
		// nums1 = [1], m = 1, nums2 = [], n = 0
		// []int{1}
		nums1 = []int{1}
		m = 1
		nums2 = []int{}
		n = 0
		merge(nums1, m, nums2, n)
		So(reflect.DeepEqual(nums1, []int{1}), ShouldBeTrue)
	})
}
