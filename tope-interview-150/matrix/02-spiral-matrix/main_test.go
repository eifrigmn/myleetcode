package main

import (
	. "github.com/smartystreets/goconvey/convey"
	"reflect"
	"testing"
)

func TestSpiralOrder(t *testing.T) {
	Convey("螺旋矩阵", t, func() {
		var matrix = [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
		var result = []int{1, 2, 3, 6, 9, 8, 7, 4, 5}
		res := spiralOrder(matrix)
		So(reflect.DeepEqual(res, result), ShouldBeTrue)

		matrix = [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}}
		result = []int{1, 2, 3, 4, 8, 12, 11, 10, 9, 5, 6, 7}
		res = spiralOrder(matrix)
		So(reflect.DeepEqual(res, result), ShouldBeTrue)

		matrix = [][]int{{7}, {9}, {6}}
		result = []int{7, 9, 6}
		res = spiralOrder(matrix)
		So(reflect.DeepEqual(res, result), ShouldBeTrue)
	})
}
