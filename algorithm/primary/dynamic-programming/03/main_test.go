package _03

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestMaxSubArray(t *testing.T) {
	Convey("最大子序和", t, func() {
		// [-2,1,-3,4,-1,2,1,-5,4] -> 6
		nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
		So(maxSubArray(nums), ShouldEqual, 6)
		// [4,1,-3,-2,-1,2,1,-5,-4] -> 5
		nums = []int{4, 1, -3, -2, -1, 2, 1, -5, -4}
		So(maxSubArray(nums), ShouldEqual, 5)
		// [1] -> 1
		nums = []int{1}
		So(maxSubArray(nums), ShouldEqual, 1)
	})
}
