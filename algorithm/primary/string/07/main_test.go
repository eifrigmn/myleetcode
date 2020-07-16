package _07

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestStrStr(t *testing.T) {
	Convey("实现 strStr()", t, func() {
		var haystack, needle string
		// haystack = "hello", needle = "ll"
		haystack = "hello"
		needle = "ll"
		So(strStr(haystack, needle), ShouldEqual, 2)
		// haystack = "aaaaa", needle = "bba"
		haystack = "aaaaa"
		needle = "bba"
		So(strStr(haystack, needle), ShouldEqual, -1)
	})
}
