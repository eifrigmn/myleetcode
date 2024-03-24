package main

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestMaxProfit(t *testing.T) {
	Convey("跳跃游戏", t, func() {
		var nums = []int{2, 3, 1, 1, 4}
		So(canJump(nums), ShouldBeTrue)

		nums = []int{3, 2, 1, 0, 4}
		So(canJump(nums), ShouldBeFalse)

		nums = []int{0}
		So(canJump(nums), ShouldBeTrue)

		nums = []int{0, 1}
		So(canJump(nums), ShouldBeFalse)

		nums = []int{1, 1, 1, 0}
		So(canJump(nums), ShouldBeTrue)

		nums = []int{2, 0}
		So(canJump(nums), ShouldBeTrue)

		nums = []int{2, 0, 2}
		So(canJump(nums), ShouldBeTrue)

		nums = []int{1, 0, 2}
		So(canJump(nums), ShouldBeFalse)

		nums = []int{1, 0, 1, 0}
		So(canJump(nums), ShouldBeFalse)
	})
}
