# 动态规划法

具体到某一天，持仓状态只能有两种：持有 (a)、未持有 (b)

以这两种持仓状态分别计算当天的累积利润，最终取这两种利润的最大值 max(a,b)即为最大利润

~~~
maxProfit = max(a,b)
~~~

+ 持有 (a)

某天 (i) 持有股票，可能存在两种动作：

1. 当天买入，说明前一天未持有股票，则此时的累积利润为：

~~~
profit1 = b(i-1) - price(i)
~~~

2. 当天未买入，那么当天持有的股票，应该来自于前一天持有的股票

~~~
profit2 = a(i-1)
~~~

由于是取利润的最大值，因此，这种情况下的当天累积利润公式为：

~~~
a = max(profit1, profit2)
即
a(i) = max(b(i-1)-price(i), a(i-1))
~~~

+ 未持有(b)

某天 (i)未持有股票，可能存在两种动作：

1. 当天未买入，那么当天的累积利润来源于前一天未持有股票的利润

~~~
profit1 = b(i-1)
~~~

2. 当天卖出，说明前一天持有股票

~~~
profit2 = a(i-1)+price(i)
~~~

同样的，当天未持有股票的累积利润公式为：

~~~
b(i) = max(b(i-1), a(i-1)+price(i))
~~~

## 边界条件

第一天，用户持有股票的累积利润为：

~~~
a(0) = -price(0)
~~~

未持有股票的累积利润为：

~~~
b(0) = 0
~~~

## 代码

~~~go
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
~~~

