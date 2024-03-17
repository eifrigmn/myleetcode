package main

import (
	. "github.com/smartystreets/goconvey/convey"
	"reflect"
	"testing"
)

func TestRemoveDuplicates(t *testing.T) {
	Convey("删除有序数组中的重复项 II", t, func() {
		var nums = []int{1, 1, 1, 2, 2, 3}
		res := removeDuplicates(nums)
		So(res, ShouldEqual, 5)
		So(reflect.DeepEqual(nums[:res], []int{1, 1, 2, 2, 3}), ShouldBeTrue)

		nums = []int{0, 0, 1, 1, 1, 1, 2, 3, 3}
		res = removeDuplicates(nums)
		So(res, ShouldEqual, 7)
		So(reflect.DeepEqual(nums[:res], []int{0, 0, 1, 1, 2, 3, 3}), ShouldBeTrue)

		nums = []int{1, 1, 1, 2, 2, 2, 3, 3}
		res = removeDuplicates(nums)
		So(res, ShouldEqual, 6)
		So(reflect.DeepEqual(nums[:res], []int{1, 1, 2, 2, 3, 3}), ShouldBeTrue)

		nums = []int{0, 1, 2, 2, 2, 2, 2, 3, 4, 4, 4}
		res = removeDuplicates(nums)
		So(res, ShouldEqual, 7)
		So(reflect.DeepEqual(nums[:res], []int{0, 1, 2, 2, 3, 4, 4}), ShouldBeTrue)

		nums = []int{1}
		res = removeDuplicates(nums)
		So(res, ShouldEqual, 1)
		So(reflect.DeepEqual(nums[:res], []int{1}), ShouldBeTrue)

		nums = []int{1, 2}
		res = removeDuplicates(nums)
		So(res, ShouldEqual, 2)
		So(reflect.DeepEqual(nums[:res], []int{1, 2}), ShouldBeTrue)

		nums = []int{1, 1}
		res = removeDuplicates(nums)
		So(res, ShouldEqual, 2)
		So(reflect.DeepEqual(nums[:res], []int{1, 1}), ShouldBeTrue)
	})
}
