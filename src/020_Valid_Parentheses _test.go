package src

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestIsValid(t *testing.T) {
	Convey("有效的括号", t, func() {
		s1 := "(("
		So(isValid(s1), ShouldBeFalse)
		s2 := "(){}[]"
		So(isValid(s2), ShouldBeTrue)
		s3 := "(]"
		So(isValid(s3), ShouldBeFalse)
		s4 := "([)]"
		So(isValid(s4), ShouldBeFalse)
		s5 := "{[}"
		So(isValid(s5), ShouldBeFalse)
		s6 := "(([]){})"
		So(isValid(s6), ShouldBeTrue)
		s7 := "[({(())}[()])]"
		So(isValid(s7), ShouldBeTrue)
	})
}
