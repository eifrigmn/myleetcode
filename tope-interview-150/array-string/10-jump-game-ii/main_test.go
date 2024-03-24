package main

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestMaxProfit(t *testing.T) {
	Convey("跳跃游戏 II", t, func() {
		var nums = []int{2, 3, 1, 1, 4}
		minSteps := jump(nums)
		So(minSteps, ShouldEqual, 2)

		nums = []int{2, 3, 0, 1, 4}
		minSteps = jump(nums)
		So(minSteps, ShouldEqual, 2)

		nums = []int{0}
		minSteps = jump(nums)
		So(minSteps, ShouldEqual, 0)

		nums = []int{1, 1}
		minSteps = jump(nums)
		So(minSteps, ShouldEqual, 1)

		nums = []int{1, 1, 1, 0}
		minSteps = jump(nums)
		So(minSteps, ShouldEqual, 3)

		nums = []int{2, 0}
		minSteps = jump(nums)
		So(minSteps, ShouldEqual, 1)

		nums = []int{2, 3, 1}
		minSteps = jump(nums)
		So(minSteps, ShouldEqual, 1)

		nums = []int{2, 1, 1, 1, 1}
		minSteps = jump(nums)
		So(minSteps, ShouldEqual, 3)

		nums = []int{1, 1, 3, 1, 1}
		minSteps = jump(nums)
		So(minSteps, ShouldEqual, 3)

		nums = []int{1, 2, 3, 4, 5}
		minSteps = jump(nums)
		So(minSteps, ShouldEqual, 3)

		nums = []int{3, 4, 3, 2, 5, 4, 3}
		minSteps = jump(nums)
		So(minSteps, ShouldEqual, 3)

		nums = []int{7, 0, 9, 6, 9, 6, 1, 7, 9, 0, 1, 2, 9, 0, 3}
		minSteps = jump(nums)
		So(minSteps, ShouldEqual, 2)
	})
}
