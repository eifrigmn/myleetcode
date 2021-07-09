package _04

import (
	"fmt"
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

func TestArr(t *testing.T) {
	var s1 = []int{1, 2, 3, 4}
	s1 = append(s1, 5)
	s2 := s1[:]
	s2 = append(s2, 8, 9)
	s1 = append(s1, 6, 7)
	fmt.Print(s1, s2)
}
