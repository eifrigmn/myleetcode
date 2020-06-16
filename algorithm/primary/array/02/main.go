package _02

func maxProfit(prices []int) int {
	var result, lft, rgt int
	for i := 0; i<len(prices)-1;i++{
		if prices[i] > prices[i+1] {
			result += prices[rgt]-prices[lft]
			lft = i+1
			rgt = i+1
		} else {
			rgt = i+1
		}
	}
	result += prices[rgt]-prices[lft]
	return result
}
