package _0003

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestAddTwoNumbers(t *testing.T) {
	Convey("lengthOfLongestSubstring", t, func() {
		s := "abcabcbb"
		l := lengthOfLongestSubstring(s)
		So(l, ShouldEqual, 3)

		s = "bbbbb"
		l = lengthOfLongestSubstring(s)
		So(l, ShouldEqual, 1)

		s = "pwwkew"
		l = lengthOfLongestSubstring(s)
		So(l, ShouldEqual, 3)

		s = " "
		l = lengthOfLongestSubstring(s)
		So(l, ShouldEqual, 1)

		s = ""
		l = lengthOfLongestSubstring(s)
		So(l, ShouldEqual, 0)

		s = "au"
		l = lengthOfLongestSubstring(s)
		So(l, ShouldEqual, 2)

		s = "aab"
		l = lengthOfLongestSubstring(s)
		So(l, ShouldEqual, 2)

		s = "abba"
		l = lengthOfLongestSubstring(s)
		So(l, ShouldEqual, 2)
	})
}
