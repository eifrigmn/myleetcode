package _01

import (
	. "github.com/smartystreets/goconvey/convey"
	"reflect"
	"testing"
)

func TestRemoveDuplicates(t *testing.T) {
	Convey("removeDuplicates", t, func() {
		// [1,1,2]
		nums1 := []int{1,1,2}
		l1 := removeDuplicates(nums1)
		So(l1, ShouldEqual, 2)
		So(reflect.DeepEqual(nums1[:l1], []int{1,2}), ShouldBeTrue)
		// [0,0,1,1,1,2,2,3,3,4]
		nums2 := []int{0,0,1,1,1,2,2,3,3,4}
		l2 := removeDuplicates(nums2)
		So(l2, ShouldEqual, 5)
		So(reflect.DeepEqual(nums2[:l2], []int{0,1,2,3,4}), ShouldBeTrue)
	})
}

