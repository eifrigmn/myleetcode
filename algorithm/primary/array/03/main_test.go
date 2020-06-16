package _03

import (
	. "github.com/smartystreets/goconvey/convey"
	"reflect"
	"testing"
)

func TestRotate(t *testing.T) {
	Convey("rotate", t, func() {
		// [1,2,3,4,5,6,7], k = 3
		var nums = []int{1,2,3,4,5,6,7}
		var k = 3
		rotate1(nums, k)
		So(reflect.DeepEqual(nums, []int{5,6,7,1,2,3,4}), ShouldBeTrue)
		// [-1,-100,3,99], k = 2
		nums = []int{-1,-100,3,99}
		k = 2
		rotate(nums, k)
		So(reflect.DeepEqual(nums, []int{3,99,-1,-100}), ShouldBeTrue)
		// [1,2], k = 3
		nums = []int{1,2}
		k = 3
		rotate(nums, k)
		So(reflect.DeepEqual(nums, []int{2,1}), ShouldBeTrue)

	})
}
