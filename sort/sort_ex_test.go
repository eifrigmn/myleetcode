package sort_test

import (
	"myleetcode/sort"
	"reflect"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

var (
	nums       = []int{2, 3, 8, 7, 1, 2, 2, 2, 7, 3, 9, 8, 2, 1, 4, 2, 4, 6, 9, 2}
	sortedNums = []int{1, 1, 2, 2, 2, 2, 2, 2, 2, 3, 3, 4, 4, 6, 7, 7, 8, 8, 9, 9}
)

func TestBubbleSort(t *testing.T) {
	Convey("冒泡排序", t, func() {
		sort.BubbleSort(nums)
		So(reflect.DeepEqual(nums, sortedNums), ShouldBeTrue)
	})
}

func TestInsertionSort(t *testing.T) {
	Convey("插入排序", t, func() {
		nums = []int{4, 1, 3, 2}
		sort.InsertSort(nums)
		So(reflect.DeepEqual(nums, sortedNums), ShouldBeTrue)
	})
}

//func TestSelectionSort(t *testing.T) {
//	Convey("选择排序", t, func() {
//		selectionSort(nums)
//		So(reflect.DeepEqual(nums, sortedNums), ShouldBeTrue)
//	})
//}
//
func TestMergeSort(t *testing.T) {
	Convey("归并排序", t, func() {
		result := sort.MergeSort(nums)
		So(reflect.DeepEqual(result, sortedNums), ShouldBeTrue)
	})
}

func TestQuickSort(t *testing.T) {
	Convey("快速排序", t, func() {
		sort.Quicksort(nums)
		So(reflect.DeepEqual(nums, sortedNums), ShouldBeTrue)
	})
}
