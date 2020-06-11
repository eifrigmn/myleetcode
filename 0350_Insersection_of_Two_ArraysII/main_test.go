package _0350

import (
	"reflect"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestInsersect(t *testing.T) {
	Convey("两个数组的交集", t, func() {
		// nums1 = [1,2,2,1]
		// nums2 = [2,2]
		nums1 := []int{1, 2, 2, 1}
		nums2 := []int{2, 2}
		inums := intersect1(nums1, nums2)
		So(reflect.DeepEqual(inums, []int{2, 2}), ShouldBeTrue)
		// nums1 = [4,9,5]
		// nums2 = [9,4,9,8,4]
		nums1 = []int{4,9,5}
		nums2 = []int{9,4,9,8,4}
		inums = intersect1(nums1, nums2)
		So(reflect.DeepEqual(inums, []int{4,9}), ShouldBeTrue)
	})
}
