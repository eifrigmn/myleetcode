package sort_ex

import (
	"reflect"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

var (
	nums       = []int{2, 3, 8, 7, 1, 2, 2, 2, 7, 3, 9, 8, 2, 1, 4, 2, 4, 6, 9, 2}
	sortedNums = []int{1, 1, 2, 2, 2, 2, 2, 2, 2, 3, 3, 4, 4, 6, 7, 7, 8, 8, 9, 9}
)

func TestBubbule(t *testing.T) {
	Convey("冒泡排序", t, func() {
		bubble(nums)
		So(reflect.DeepEqual(nums, sortedNums), ShouldBeTrue)
	})
}

func TestSelection(t *testing.T) {
	Convey("选择排序", t, func() {
		selection(nums)
		So(reflect.DeepEqual(nums, sortedNums), ShouldBeTrue)
	})
}

func TestInsert(t *testing.T) {
	Convey("插入排序", t, func() {
		insert(nums)
		So(reflect.DeepEqual(nums, sortedNums), ShouldBeTrue)
	})
}

func TestMerge(t *testing.T) {
	Convey("归并排序", t, func() {
		nums = merge(nums)
		So(reflect.DeepEqual(nums, sortedNums), ShouldBeTrue)
	})
}

func TestQuickSort(t *testing.T) {
	Convey("快排", t, func() {
		quickSort(nums)
		So(reflect.DeepEqual(nums, sortedNums), ShouldBeTrue)
	})
}
