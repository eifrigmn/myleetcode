package _017

import (
	. "github.com/smartystreets/goconvey/convey"
	"reflect"
	"testing"
)

func Test_threeSumClosest(t *testing.T) {
	Convey("isPalindrome", t, func() {
		So(reflect.DeepEqual(letterCombinations("23"),[]string{"ad","ae","af","bd","be","bf","cd","ce","cf"}), ShouldBeTrue)
		So(reflect.DeepEqual(letterCombinations(""),[]string{}), ShouldBeTrue)
	})
}
