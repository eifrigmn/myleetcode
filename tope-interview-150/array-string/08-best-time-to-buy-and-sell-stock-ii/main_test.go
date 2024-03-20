package main

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestMaxProfit(t *testing.T) {
	Convey("买卖股票的最佳时机 II", t, func() {
		var prices = []int{7, 1, 5, 3, 6, 4}
		profit := _maxProfit1(prices)
		So(profit, ShouldEqual, 7)

		prices = []int{77, 99, 1, 3, 6, 4}
		profit = _maxProfit1(prices)
		So(profit, ShouldEqual, 27)

		prices = []int{1, 2, 3, 4, 5}
		profit = _maxProfit1(prices)
		So(profit, ShouldEqual, 4)

		prices = []int{1, 9, 6, 9, 1, 7, 1, 1, 5, 9, 9, 9}
		profit = _maxProfit1(prices)
		So(profit, ShouldEqual, 25)
	})
}
