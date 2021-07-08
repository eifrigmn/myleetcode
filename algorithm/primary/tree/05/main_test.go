package _05

import (
	"myleetcode/util"
	"reflect"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

var null = util.NULL

func TestSortedArrayToBST(t *testing.T) {
	Convey("将有序数组转换为二叉搜索树", t, func() {
		//        0				0
		//		 / \		   / \
		// 		-3  9         -10 5
		//		/   /          \  \
		// 	   -10  5          -3  9
		nums := []int{-10, -3, 0, 5, 9}
		root := sortedArrayToBST(nums)
		So(reflect.DeepEqual(util.TreeToInts(root), []int{0, -3, 9, -10, null, 5}) || reflect.DeepEqual(util.TreeToInts(root), []int{0, -10, 5, null, -3, null, 9}), ShouldBeTrue)
		//      3  		 		1
		//      /   			\
		//      1    			3
		nums = []int{1, 3}
		root = sortedArrayToBST(nums)
		So(reflect.DeepEqual(util.TreeToInts(root), []int{3, 1}) || reflect.DeepEqual(util.TreeToInts(root), []int{1, null, 3}), ShouldBeTrue)
	})
}
