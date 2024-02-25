package _022

import (
	. "github.com/smartystreets/goconvey/convey"
	"reflect"
	"testing"
)

func Test_generateParenthesis(t *testing.T) {
	Convey("generateParenthesis", t, func() {
		// n=0
		n := 1
		result := generateParenthesis(n)
		target := []string{"()"}
		So(reflect.DeepEqual(result, target), ShouldBeTrue)
		// 3
		n = 3
		target = []string{"()()()", "()(())", "(())()", "(()())", "((()))"}
		result = generateParenthesis(n)
		t.Log(result)
		So(reflect.DeepEqual(result, target), ShouldBeTrue)
	})
}
