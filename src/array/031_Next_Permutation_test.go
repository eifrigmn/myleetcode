package array

import (
	. "github.com/smartystreets/goconvey/convey"
	"reflect"
	"testing"
)

func TestNextPermutation(t *testing.T) {
	Convey("next permutation", t, func() {
		var (
			nums1 = []int{1, 5, 8, 4, 7, 6, 5, 3, 1}
			nums2 = []int{1,2,5,8,7}
			nums3 = []int{3,2,1}
			nums4 = []int{1,2,3}
			nums5 = []int{1,1,5}
			nums6 = []int{1,2}
		)
		nextPermutation(nums1)
		So(reflect.DeepEqual(nums1, []int{1, 5, 8, 5, 1, 3, 4, 6, 7}), ShouldBeTrue)
		nextPermutation(nums2)
		So(reflect.DeepEqual(nums2, []int{1, 2, 7, 5, 8}), ShouldBeTrue)
		nextPermutation(nums3)
		So(reflect.DeepEqual(nums3, []int{1, 2, 3}), ShouldBeTrue)
		nextPermutation(nums4)
		So(reflect.DeepEqual(nums4, []int{1, 3, 2}), ShouldBeTrue)
		nextPermutation(nums5)
		So(reflect.DeepEqual(nums5, []int{1, 5, 1}), ShouldBeTrue)
		nextPermutation(nums6)
		So(reflect.DeepEqual(nums6, []int{2, 1}), ShouldBeTrue)
	})
}
