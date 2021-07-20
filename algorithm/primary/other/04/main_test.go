package _04

import (
	"reflect"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestGenerate(t *testing.T) {
	Convey("杨辉三角形", t, func() {
		numRows := 5
		result := [][]int{
			{1},
			{1, 1},
			{1, 2, 1},
			{1, 3, 3, 1},
			{1, 4, 6, 4, 1},
		}
		So(reflect.DeepEqual(generate(numRows), result), ShouldBeTrue)
	})
}
