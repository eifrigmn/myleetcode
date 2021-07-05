package _09

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestLongestCommonPrefix(t *testing.T) {
	Convey("最长公共前缀", t, func() {
		var strs []string
		// ["flower","flow","flight"] -> "fl"
		strs = []string{"flower", "flow", "flight"}
		So(longestCommonPrefix1(strs), ShouldEqual, "fl")
		// ["dog","racecar","car"] -> ""
		strs = []string{"dog", "racecar", "car"}
		So(longestCommonPrefix1(strs), ShouldEqual, "")
	})
}
