package _01

import (
	"reflect"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestFizzBuzz(t *testing.T) {
	Convey("fizzBuzz", t, func() {
		// n = 15
		// [
		//		"1",
		//		"2",
		//		"Fizz",
		//		"4",
		//		"Buzz",
		//		"Fizz",
		//		"7",
		//		"8",
		//		"Fizz",
		//		"Buzz",
		//		"11",
		//		"Fizz",
		//		"13",
		//		"14",
		//		"FizzBuzz"
		//	]
		var n = 15
		var result = []string{"1", "2", "Fizz", "4", "Buzz", "Fizz", "7", "8", "Fizz", "Buzz", "11", "Fizz", "13", "14", "FizzBuzz"}
		So(reflect.DeepEqual(fizzBuzz1(n), result), ShouldBeTrue)
	})
}
