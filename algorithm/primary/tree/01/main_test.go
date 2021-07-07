package _01

import (
	"myleetcode/util"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

var null = util.NULL

func TestReverseString(t *testing.T) {
	Convey("二叉树最大深度", t, func() {
		//     3
		//    / \
		//	 9  20
		//     /  \
		//    15   7
		root := util.IntsToTreeNode([]int{3, 9, 20, null, null, 15, 7})
		So(maxDepth(root), ShouldEqual, 3)
	})
}
