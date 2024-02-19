package _0003

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestAddTwoNumbers(t *testing.T) {
	Convey("lengthOfLongestSubstring", t, func() {
		s := "abcabcbb"
		l := _lengthOfLongestSubstring(s)
		So(l, ShouldEqual, 3)

		s = "bbbbb"
		l = _lengthOfLongestSubstring(s)
		So(l, ShouldEqual, 1)

		s = "pwwkew"
		l = _lengthOfLongestSubstring(s)
		So(l, ShouldEqual, 3)

		s = " "
		l = _lengthOfLongestSubstring(s)
		So(l, ShouldEqual, 1)

		s = ""
		l = _lengthOfLongestSubstring(s)
		So(l, ShouldEqual, 0)

		s = "au"
		l = _lengthOfLongestSubstring(s)
		So(l, ShouldEqual, 2)

		s = "aab"
		l = _lengthOfLongestSubstring(s)
		So(l, ShouldEqual, 2)

		s = "abba"
		l = _lengthOfLongestSubstring(s)
		So(l, ShouldEqual, 2)

		s = "dvdf"
		l = _lengthOfLongestSubstring(s)
		So(l, ShouldEqual, 3)
	})
}
