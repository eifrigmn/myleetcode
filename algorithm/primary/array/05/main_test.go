package _05

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestSingleNumber(t *testing.T) {
	Convey("只出现一次的数字", t, func(){
		// [2,2,1]
		nums := []int{2,2,1}
		So(singleNumber(nums), ShouldEqual, 1)
		// [4,1,2,1,2]
		nums = []int{4,1,2,1,2}
		So(singleNumber(nums), ShouldEqual, 4)
	})
}
