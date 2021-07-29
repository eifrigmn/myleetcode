package _01

import "math"

func climbStairs_math(n int) int {
	m1 := math.Pow((1+math.Sqrt(5))/2, float64(n+1)) - math.Pow((1-math.Sqrt(5))/2, float64(n+1))
	return int(m1 / math.Sqrt(5))
}
