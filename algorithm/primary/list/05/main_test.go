package _05

import (
	"myleetcode/util"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestIsPalindrome(t *testing.T) {
	Convey("回文链表", t, func() {
		// 1->2
		// false
		var l *ListNode
		l = util.IntsToSList([]int{1, 2})
		So(isPalindrome1(l), ShouldBeFalse)
		// 1->2->2->1
		// true
		l = util.IntsToSList([]int{1, 2, 2, 1})
		So(isPalindrome1(l), ShouldBeTrue)
		// 1->1->2->1
		// false
		l = util.IntsToSList([]int{1, 1, 2, 1})
		So(isPalindrome1(l), ShouldBeFalse)
	})
}
