package _0136

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestSingleNumber(t *testing.T) {
	Convey("single number", t, func() {
		var (
			nums1 = []int{2, 2, 1}
			nums2 = []int{4, 1, 2, 1, 2}
		)
		So(singleNumber(nums1), ShouldEqual, 1)
		So(singleNumber(nums2), ShouldEqual, 4)
	})
}
