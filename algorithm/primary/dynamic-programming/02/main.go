package _02

// 寻找最长的升序区间
func maxProfit(prices []int) int {
	if len(prices) < 2 {
		return 0
	}

	var min, maxPro int = prices[0], 0
	for i := 1; i < len(prices); i++ {
		if prices[i] < min {
			min = prices[i]
		}

		if prices[i]-min > maxPro {
			maxPro = prices[i] - min
		}
	}
	return maxPro
}
