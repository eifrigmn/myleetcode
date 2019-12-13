package src

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestRemoveElement(t *testing.T) {
	Convey("Remove Element", t, func() {
		var nums = []int{3, 2, 2, 3}
		var val = 3
		So(removeElement(nums, val), ShouldEqual, 2)
		nums = []int{0,1,2,2,3,0,4,2}
		val = 2
		So(removeElement(nums, val), ShouldEqual, 5)
		nums = []int{}
		val = 1
		So(removeElement(nums, val), ShouldEqual, 0)
		nums = []int{2,2}
		val = 2
		So(removeElement(nums, val), ShouldEqual, 0)
	})
}