package main

func maxProfit(prices []int) int {
	// 把价格序列分割成多个升序子数组，求每个子数组的最大利润之和
	var slow, profit int
	var fast = 1
	for fast < len(prices) {
		if prices[fast] < prices[fast-1] {
			profit += prices[fast-1] - prices[slow]
			slow = fast
		} else if fast == len(prices)-1 && prices[fast] >= prices[fast-1] {
			profit += prices[fast] - prices[slow]
		}
		fast++
	}
	return profit
}

func _maxProfit(prices []int) int {
	sum := 0
	for i := 1; i < len(prices); i++ {
		profit := prices[i] - prices[i-1]
		if profit > 0 {
			sum += profit
		}
	}
	return sum
}

// 动态规划
// dp0数组存储的是第i天未持有股票时的最大利润
// dp1数组存储的时第i天持有股票时的最大利润
// 那么第i天未持有股票时的最大利润，为下述两种情况取利润较大的一个：
// 1. 第i-1天未持有股票
// 2. 第i-1天持有股票，在第i天卖出股票
// 即：dp0[i] = max(dp0[i-1], dp1[i-1]+prices[i])
// 状态转移方程中涉及dp1，那么第i天持有股票时的最大利润为：
// dp1[i] = max(dp1[i-1], dp0[i-1]-prices[i])
// 初始状态为：
// dp0[0] = 0
// dp1[0] = -prices[0]
func _maxProfit1(prices []int) int {
	// 第i天未持有股票时的最大利润
	var dp0 = make([]int, len(prices))
	// 第i天持有股票时的最大利润
	var dp1 = make([]int, len(prices))
	dp0[0] = 0
	dp1[0] = -prices[0]
	for i := 1; i < len(prices); i++ {
		dp0[i] = maxInt(dp0[i-1], dp1[i-1]+prices[i])
		dp1[i] = maxInt(dp1[i-1], dp0[i-1]-prices[i])
	}
	return dp0[len(prices)-1]
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}
