package _08

import (
	. "github.com/smartystreets/goconvey/convey"
	"reflect"
	"testing"
)

func TestMoveZeroes(t *testing.T)  {
	Convey("移动零", t, func() {
		nums := []int{0,1,0,3,12}
		moveZeroes(nums)
		So(reflect.DeepEqual(nums, []int{1,3,12,0,0}), ShouldBeTrue)
		nums = []int{1,2,3,4,5}
		moveZeroes(nums)
		So(reflect.DeepEqual(nums, []int{1,2,3,4,5}), ShouldBeTrue)
	})
}
