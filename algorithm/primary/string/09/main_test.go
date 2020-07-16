package _09

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestLongestCommonPrefix(t *testing.T) {
	Convey("最长公共前缀", t, func() {
		var strs []string
		// ["flower","flow","flight"] -> "fl"
		strs = []string{"flower","flow","flight"}
		So(longestCommonPrefix(strs), ShouldEqual, "fl")
		// ["dog","racecar","car"] -> ""
		strs = []string{"dog","racecar","car"}
		So(longestCommonPrefix(strs), ShouldEqual, "")
	})
}
