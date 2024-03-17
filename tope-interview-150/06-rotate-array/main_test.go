package main

import (
	. "github.com/smartystreets/goconvey/convey"
	"reflect"
	"testing"
)

func TestRotate(t *testing.T) {
	Convey("轮转数组", t, func() {
		var nums = []int{1, 2, 3, 4, 5, 6, 7}
		var k = 3
		rotate(nums, k)
		So(reflect.DeepEqual(nums, []int{5, 6, 7, 1, 2, 3, 4}), ShouldBeTrue)

		nums = []int{-1}
		k = 2
		rotate(nums, k)
		So(reflect.DeepEqual(nums, []int{-1}), ShouldBeTrue)
	})
}
