package _02

import (
	"myleetcode/util"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

var null = util.NULL

func TestReverseString(t *testing.T) {
	Convey("验证二叉搜索树", t, func() {
		//     2
		//    / \
		//	 1   3
		// out: true
		root := util.IntsToTreeNode([]int{2, 1, 3})
		So(isValidBST(root), ShouldBeTrue)
		//     5
		//	  / \
		//   1   4
		// 		/ \
		// 	   3   6
		// out: false
		root = util.IntsToTreeNode([]int{5, 1, 4, null, null, 3, 6})
		So(isValidBST(root), ShouldBeFalse)
		//   2
		//  / \
		// 2   2
		// out: false
		root = util.IntsToTreeNode([]int{2, 2, 2})
		So(isValidBST(root), ShouldBeFalse)
		//     5
		//	  / \
		//   4   6
		// 		/ \
		// 	   3   7
		// out: false
		root = util.IntsToTreeNode([]int{5, 4, 6, null, null, 3, 7})
		So(isValidBST(root), ShouldBeFalse)
	})
}
