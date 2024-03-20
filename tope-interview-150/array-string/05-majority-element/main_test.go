package main

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestMajorityElement(t *testing.T) {
	Convey("多数元素", t, func() {
		var nums = []int{3, 2, 3}
		element := _majorityElement(nums)
		So(element, ShouldEqual, 3)

		nums = []int{2, 2, 1, 1, 1, 2, 2}
		element = _majorityElement(nums)
		So(element, ShouldEqual, 2)

		nums = []int{0, 0, 0, 0}
		element = _majorityElement(nums)
		So(element, ShouldEqual, 0)
	})
}
