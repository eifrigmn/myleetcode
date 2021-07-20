package _06

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestMissingNumber(t *testing.T) {
	Convey("缺失数字", t, func() {
		// nums = [3,0,1] -> 2
		nums := []int{3, 0, 1}
		So(missingNumber(nums), ShouldEqual, 2)
		// nums = [0,1] -> 2
		nums = []int{0, 1}
		So(missingNumber(nums), ShouldEqual, 2)
		// nums = [9,6,4,2,3,5,7,0,1] -> 8
		nums = []int{9, 6, 4, 2, 3, 5, 7, 0, 1}
		So(missingNumber(nums), ShouldEqual, 8)

	})
}
