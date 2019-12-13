package src

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestStrStr(t *testing.T) {
	Convey("Implement strStr()", t, func() {
		var haystack = "hello"
		var needle = "ll"
		index := strStr(haystack, needle)
		So(index, ShouldEqual, 2)

		haystack = "aaaaa"
		needle = "bba"
		index = strStr(haystack, needle)
		So(index, ShouldEqual, -1)

		haystack = "bba"
		needle = "aaaaa"
		index = strStr(haystack, needle)
		So(index, ShouldEqual, -1)

		haystack = ""
		needle = "a"
		index = strStr(haystack, needle)
		So(index, ShouldEqual, -1)
	})
}
