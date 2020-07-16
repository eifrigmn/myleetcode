package _01

import (
	. "github.com/smartystreets/goconvey/convey"
	"myleetcode/util"
	"reflect"
	"testing"
)

func TestDeleteNode(t *testing.T) {
	Convey("删除链表中的节点", t, func() {
		// head = [4,5,1,9], node = 5
		var l, node *ListNode
		var result []int
		l = util.IntsToSList([]int{4,5,1,9})
		node = l.Next
		deleteNode(node)
		result = util.SListToInts(l)
		So(reflect.DeepEqual(result, []int{4,1,9}), ShouldBeTrue)
		// head = [4,5,1,9], node = 1
		l = util.IntsToSList([]int{4,5,1,9})
		node = l.Next.Next
		deleteNode(node)
		result = util.SListToInts(l)
		So(reflect.DeepEqual(result, []int{4,5,9}), ShouldBeTrue)
	})
}
