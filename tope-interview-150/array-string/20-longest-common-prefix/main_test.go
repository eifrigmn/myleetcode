package main

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestLongestCommonPrefix(t *testing.T) {
	Convey("最长公共前缀", t, func() {
		strs := []string{"flower", "flow", "flight"}
		prefix := longestCommonPrefix(strs)
		So(prefix, ShouldEqual, "fl")

		strs = []string{"dog", "racecar", "car"}
		prefix = longestCommonPrefix(strs)
		So(prefix, ShouldEqual, "")
	})
}
