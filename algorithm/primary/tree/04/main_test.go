package _04

import (
	"myleetcode/util"
	"reflect"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

var null = util.NULL

func TestLevelOrder(t *testing.T) {
	Convey("对称二叉树", t, func() {
		//     3
		//    / \
		//   9  20
		//     /  \
		// 	  15   7
		root := util.IntsToTreeNode([]int{3, 9, 20, null, null, 15, 7})
		result := [][]int{{3}, {9, 20}, {15, 7}}
		So(reflect.DeepEqual(levelOrder(root), result), ShouldBeTrue)
	})
}
