# 贪心算法

在每一次选择中，都采取当前状态下的最好选择，从而希望导致结果为最优解。

贪心算法不同于动态规划，贪心算法对每个子问题的解决方案都做出选择，不能回退；而动态规划会保存以前的运算结果，并根据以前的运算结果，对当前进行选择。

## 解题思路：

设置两个指针`buy`和`sell`，寻找每次股票下跌的拐点，作为卖出点`sell`

`sell`与`buy`之间的差值，则为此次操作的利润

此后，`sell`指针和`buy`指针均移动到`sell`指针的后一位

~~~go
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
~~~
