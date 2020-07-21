package _05

import (
	. "github.com/smartystreets/goconvey/convey"
	"myleetcode/util"
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	Convey("回文链表", t, func() {
		// 1->2
		// false
		var l *ListNode
		l = util.IntsToSList([]int{1,2})
		So(isPalindrome(l), ShouldBeFalse)
		// 1->2->2->1
		// true
		l = util.IntsToSList([]int{1,2,2,1})
		So(isPalindrome(l), ShouldBeTrue)
	})
}
