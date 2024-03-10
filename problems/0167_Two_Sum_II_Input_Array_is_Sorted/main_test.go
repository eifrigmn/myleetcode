package _0160

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestGetIntersectionNode(t *testing.T) {
	Convey("two sum ii", t, func() {
		nums := []int{2, 7, 11, 15}
		target := 9
		result := []int{1, 2}
		So(twoSum(nums, target), ShouldEqual, result)
		nums = []int{2, 3, 4}
		target = 6
		result = []int{1, 3}
		So(twoSum(nums, target), ShouldEqual, result)
		nums = []int{-1, 0}
		target = -1
		result = []int{1, 2}
		So(twoSum(nums, target), ShouldEqual, result)
		nums = []int{5, 25, 75}
		target = 100
		result = []int{2, 3}
		So(twoSum(nums, target), ShouldEqual, result)
		nums = []int{1, 3, 4, 4}
		target = 8
		result = []int{3, 4}
		So(twoSum(nums, target), ShouldEqual, result)
	})
}
