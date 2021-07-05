package _03

import (
	"myleetcode/util"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

var null = util.NULL

func TestReverseString(t *testing.T) {
	Convey("对称二叉树", t, func() {
		//     1
		//    / \
		//   2   2
		//  / \ / \
		// 3  4 4  3
		root := util.IntsToTreeNode([]int{1, 2, 2, 3, 4, 4, 3})
		So(isSymmetric(root), ShouldBeTrue)
		//     1
		//    / \
		//   2   2
		//    \   \
		//     3    3
		// out: false
		root = util.IntsToTreeNode([]int{1, 2, 2, null, 3, null, 3})
		So(isSymmetric(root), ShouldBeFalse)
	})
}
