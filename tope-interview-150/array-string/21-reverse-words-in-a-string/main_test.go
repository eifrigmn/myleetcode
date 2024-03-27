package main

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestReverseWords(t *testing.T) {
	Convey("反转字符串中的单词", t, func() {
		s := "the sky is blue"
		rs := reverseWords(s)
		So(rs, ShouldEqual, "blue is sky the")

		s = "  hello world  "
		rs = reverseWords(s)
		So(rs, ShouldEqual, "world hello")

		s = "    "
		rs = reverseWords(s)
		So(rs, ShouldEqual, "")

		s = "a good   example"
		rs = reverseWords(s)
		So(rs, ShouldEqual, "example good a")

		s = "a "
		rs = reverseWords(s)
		So(rs, ShouldEqual, "a")

		s = "  a "
		rs = reverseWords(s)
		So(rs, ShouldEqual, "a")

		s = " asdasd df f"
		rs = reverseWords(s)
		So(rs, ShouldEqual, "f df asdasd")
	})
}
