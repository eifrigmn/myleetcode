package search

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

var nums = []int{1, 3, 4, 5, 6, 8, 8, 8, 11, 18}

func TestBinarySearchUp(t *testing.T) {
	Convey("查找第一个大于等于给定值的元素", t, func() {
		So(bSearchUp(nums, 2), ShouldEqual, 1)
		So(bSearchUp(nums, 8), ShouldEqual, 5)
		So(bSearchUp(nums, 1), ShouldEqual, 0)
		So(bSearchUp(nums, 100), ShouldEqual, -1)
	})

}

func TestBinarySearchTail(t *testing.T) {
	Convey("最后一个等于给定值的元素", t, func() {
		So(bSearchTail(nums, 8), ShouldEqual, 7)
		So(bSearchTail(nums, 9), ShouldEqual, -1)
		So(bSearchTail(nums, 1), ShouldEqual, 0)
	})
}

func TestBinarySearchHead(t *testing.T) {
	Convey("第一个等于给定值的元素", t, func() {
		So(bSearchHead(nums, 8), ShouldEqual, 5)
		So(bSearchHead(nums, 1), ShouldEqual, 0)
		So(bSearchHead(nums, 18), ShouldEqual, 9)
		So(bSearchHead(nums, 2), ShouldEqual, -1)
	})
}

func TestBinarySearchDown(t *testing.T) {
	Convey("最后一个小于等于给定值的元素", t, func() {
		So(bSearchDown(nums, 8), ShouldEqual, 7)
		So(bSearchDown(nums, 12), ShouldEqual, 8)
		So(bSearchDown(nums, 2), ShouldEqual, 0)
		So(bSearchDown(nums, 0), ShouldEqual, -1)
	})
}
