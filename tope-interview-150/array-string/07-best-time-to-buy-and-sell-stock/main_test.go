package main

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestMaxProfit(t *testing.T) {
	Convey("买卖股票的最佳时机", t, func() {
		var prices = []int{7, 1, 5, 3, 6, 4}
		profit := maxProfit(prices)
		So(profit, ShouldEqual, 5)

		prices = []int{77, 99, 1, 3, 6, 4}
		profit = maxProfit(prices)
		So(profit, ShouldEqual, 22)
	})
}
