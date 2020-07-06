package _02

import "math"

func reverse(x int) int {
	if x > math.MaxInt32 || x < math.MinInt32 {
		return 0
	}
	var isNegitave bool
	var result int
	if x < 0 {
		isNegitave = true
		x = -x
	}
	for x != 0 {
		result = result * 10 + x%10
		x = x/10
	}

	if isNegitave {
		result = -result
	}

	if result > math.MaxInt32 || result < math.MinInt32 {
		return 0
	}
	return result
}
