package _0041

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestFirstMissingPositive(t *testing.T) {
	Convey("First Missing Positive", t, func() {
		var nums = []int{1, 2, 0}
		So(firstMissingPositive(nums), ShouldEqual, 3)
		nums = []int{3, 4, -1, 1}
		So(firstMissingPositive(nums), ShouldEqual, 2)
		nums = []int{7, 8, 9, 11, 12}
		So(firstMissingPositive(nums), ShouldEqual, 1)
		nums = []int{0}
		So(firstMissingPositive(nums), ShouldEqual, 1)
		nums = []int{-1, -2}
		So(firstMissingPositive(nums), ShouldEqual, 1)
		nums = []int{0, 2, 2, 1, 1}
		So(firstMissingPositive(nums), ShouldEqual, 3)
	})

}
