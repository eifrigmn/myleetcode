package util

import (
	"reflect"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

var null = NULL

func TestIntsToTreeNode(t *testing.T) {
	Convey("测试数组转化为树", t, func() {
		//        4
		//     /     \
		//    /       \
		//    1        6
		//  /   \     /  \
		//  0   2    5    7
		//       \         \
		//        3         8
		nums := []int{4, 1, 6, 0, 2, 5, 7, null, null, null, 3, null, null, null, 8}
		tree := IntsToTreeNode(nums)
		tnums := TreeToInts(tree)
		So(reflect.DeepEqual(nums, tnums), ShouldBeTrue)
	})
}
func TestTreeToInorder(t *testing.T) {
	Convey("中序遍历", t, func() {
		nums := []int{4, 1, 6, 0, 2, 5, 7, null, null, null, 3, null, null, null, 8}
		root := IntsToTreeNode(nums)
		numsInorder := TreeToInorder1(root)
		So(reflect.DeepEqual(numsInorder, []int{0, 1, 2, 3, 4, 5, 6, 7, 8}), ShouldBeTrue)
	})
}

func TestTreeToPreorder(t *testing.T) {
	Convey("前序遍历", t, func() {
		nums := []int{4, 1, 6, 0, 2, 5, 7, null, null, null, 3, null, null, null, 8}
		root := IntsToTreeNode(nums)
		numsPre := TreeToPreorder(root)
		So(reflect.DeepEqual(numsPre, []int{4, 1, 0, 2, 3, 6, 5, 7, 8}), ShouldBeTrue)
	})
}

func TestTreeToPostorder(t *testing.T) {
	Convey("后序遍历", t, func() {
		nums := []int{4, 1, 6, 0, 2, 5, 7, null, null, null, 3, null, null, null, 8}
		root := IntsToTreeNode(nums)
		numsPostorder := TreeToPostorder(root)
		So(reflect.DeepEqual(numsPostorder, []int{0, 3, 2, 1, 5, 8, 7, 6, 4}), ShouldBeTrue)
	})
}
