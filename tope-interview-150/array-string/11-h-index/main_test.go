package main

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestMaxProfit(t *testing.T) {
	Convey("H指数", t, func() {
		var nums = []int{3, 0, 6, 1, 5}
		result := hIndex(nums)
		So(result, ShouldEqual, 3)

		nums = []int{1, 3, 1}
		result = hIndex(nums)
		So(result, ShouldEqual, 1)

		nums = []int{0}
		result = hIndex(nums)
		So(result, ShouldEqual, 0)

		nums = []int{3}
		result = hIndex(nums)
		So(result, ShouldEqual, 1)

		nums = []int{11, 15}
		result = hIndex(nums)
		So(result, ShouldEqual, 2)
	})
}
