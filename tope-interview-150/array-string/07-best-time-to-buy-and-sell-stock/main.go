package main

func maxProfit(prices []int) int {
	var cost = prices[0]
	var profit = 0
	for i := 1; i < len(prices); i++ {
		// 成本为第i天前的最小值
		if cost > prices[i] {
			cost = prices[i]
		}
		// 最大收益为i天前的最大收益
		if prices[i]-cost > profit {
			profit = prices[i] - cost
		}
	}
	return profit
}
