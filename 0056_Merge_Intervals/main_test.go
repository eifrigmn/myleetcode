package _0056

import (
	"reflect"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestMergeIntervals(t *testing.T) {
	Convey("合并区间", t, func() {
		// [[1,3],[2,6],[8,10],[15,18]] => [[1,6],[8,10],[15,18]]
		intervals := [][]int{[]int{1, 3}, []int{2, 6}, []int{8, 10}, []int{15, 18}}
		mIntervals := merge(intervals)
		So(reflect.DeepEqual(mIntervals, [][]int{[]int{1, 6}, []int{8, 10}, []int{15, 18}}), ShouldBeTrue)
		// [[1,4],[4,5]] => [[1,5]]
		intervals = [][]int{[]int{1, 4}, []int{4, 5}}
		mIntervals = merge(intervals)
		So(reflect.DeepEqual(mIntervals, [][]int{[]int{1, 5}}), ShouldBeTrue)
		// [[2,3],[2,2],[3,3],[1,3],[5,7],[2,2],[4,6]] => [[1,3],[4,7]]
		intervals = [][]int{[]int{2, 3}, []int{2, 2}, []int{3, 3}, []int{1, 3}, []int{5, 7}, []int{2, 2}, []int{4, 6}}
		mIntervals = merge(intervals)
		So(reflect.DeepEqual(mIntervals, [][]int{[]int{1, 3}, []int{4, 7}}), ShouldBeTrue)
	})
}
