package _05

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestIsValid(t *testing.T) {
	Convey("有效的括号", t, func() {
		// 输入：s = "()"
		// 输出：true
		So(isValid("()"), ShouldBeTrue)
		// 输入：s = "()[]{}"
		// 输出：true
		So(isValid("()[]{}"), ShouldBeTrue)
		// 输入：s = "(]"
		// 输出：false
		So(isValid("(]"), ShouldBeFalse)
		// 输入：s = "([)]"
		// 输出：false
		So(isValid("([)]"), ShouldBeFalse)
		// 输入：s = "{[]}"
		// 输出：true
		So(isValid("{[]}"), ShouldBeTrue)
		// 输入：s = "["
		// 输出：false
		So(isValid("["), ShouldBeFalse)
		// 输入：s = "(("
		// 输出：false
		So(isValid("(("), ShouldBeFalse)
	})
}
