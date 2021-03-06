package _0234

import (
	"myleetcode/util"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestIsPalindrome(t *testing.T) {
	Convey("is palindrome", t, func() {
		//1->2
		l1 := util.IntsToSList([]int{1, 2})
		So(isPalindrome2(l1), ShouldBeFalse)
		// 1->2->2->1
		l2 := util.IntsToSList([]int{1, 2, 2, 1})
		So(isPalindrome2(l2), ShouldBeTrue)
		// 空链表
		l3 := util.IntsToSList([]int{})
		So(isPalindrome2(l3), ShouldBeTrue)
		// 1->1->2->1
		l4 := util.IntsToSList([]int{1, 1, 2, 1})
		So(isPalindrome2(l4), ShouldBeFalse)
	})
}
