package _02

// 贪心算法
func maxProfit_greedy(prices []int) int {
	var sell, buy, profits int
	for i := 0; i < len(prices)-1; i++ {
		if prices[i] > prices[i+1] {
			profits += prices[sell] - prices[buy]
			buy = i + 1
			sell = i + 1
		} else {
			sell++
		}
	}
	profits += prices[sell] - prices[buy]
	return profits
}
