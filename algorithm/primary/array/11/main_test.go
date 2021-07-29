package _11

import (
	"reflect"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestRotate(t *testing.T) {
	Convey("旋转图像", t, func() {
		matrix := [][]int{
			{1, 2, 3},
			{4, 5, 6},
			{7, 8, 9},
		}
		rotate1(matrix)
		So(reflect.DeepEqual(matrix, [][]int{{7, 4, 1}, {8, 5, 2}, {9, 6, 3}}), ShouldBeTrue)
		matrix = [][]int{
			{5, 1, 9, 11},
			{2, 4, 8, 10},
			{13, 3, 6, 7},
			{15, 14, 12, 16},
		}
		rotate1(matrix)
		So(reflect.DeepEqual(matrix, [][]int{{15, 13, 2, 5}, {14, 3, 4, 1}, {12, 6, 8, 9}, {16, 7, 10, 11}}), ShouldBeTrue)
	})

}
