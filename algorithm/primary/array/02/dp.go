package _02

func maxProfit_dp(prices []int) int {
	if len(prices) < 2 {
		return 0
	}
	// a(i) = max(b(i-1)-price(i), a(i-1))
	// b(i) = max(b(i-1), a(i-1)+price(i))
	a := -prices[0]
	b := 0
	for i := 1; i < len(prices); i++ {
		prva := a
		prvb := b
		a = max(prvb-prices[i], prva)
		b = max(prvb, prva+prices[i])

	}
	return max(a, b)
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
