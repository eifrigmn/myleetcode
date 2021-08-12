package _0003

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestAddTwoNumbers(t *testing.T) {
	Convey("lengthOfLongestSubstring", t, func() {
		s := "abcabcbb"
		l := lengthOfLongestSubstring1(s)
		So(l, ShouldEqual, 3)

		s = "bbbbb"
		l = lengthOfLongestSubstring1(s)
		So(l, ShouldEqual, 1)

		s = "pwwkew"
		l = lengthOfLongestSubstring1(s)
		So(l, ShouldEqual, 3)

		s = " "
		l = lengthOfLongestSubstring1(s)
		So(l, ShouldEqual, 1)

		s = ""
		l = lengthOfLongestSubstring1(s)
		So(l, ShouldEqual, 0)

		s = "au"
		l = lengthOfLongestSubstring1(s)
		So(l, ShouldEqual, 2)

		s = "aab"
		l = lengthOfLongestSubstring1(s)
		So(l, ShouldEqual, 2)

		s = "abba"
		l = lengthOfLongestSubstring1(s)
		So(l, ShouldEqual, 2)

		s = "dvdf"
		l = lengthOfLongestSubstring1(s)
		So(l, ShouldEqual, 3)
	})
}
