package main

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestRomanToInt(t *testing.T) {
	Convey("最后一个单词的长度", t, func() {
		s := "Hello World"
		cnt := lengthOfLastWord(s)
		So(cnt, ShouldEqual, 5)

		s = "   fly me   to   the moon  "
		cnt = lengthOfLastWord(s)
		So(cnt, ShouldEqual, 4)

		s = "luffy is still joyboy"
		cnt = lengthOfLastWord(s)
		So(cnt, ShouldEqual, 6)
	})
}
