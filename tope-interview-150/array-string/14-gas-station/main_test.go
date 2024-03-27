package main

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestCanCompleteCircuit(t *testing.T) {
	Convey("加油站", t, func() {
		gas := []int{1, 2, 3, 4, 5}
		cost := []int{3, 4, 5, 1, 2}
		result := canCompleteCircuit(gas, cost)
		So(result, ShouldEqual, 3)

		gas = []int{2, 3, 4}
		cost = []int{3, 4, 3}
		result = canCompleteCircuit(gas, cost)
		So(result, ShouldEqual, -1)

	})
}
