package _01

func climbStairs_dp(n int) int {
	if n < 3 {
		return n
	}

	return f(n, 1, 1)
}

func f(n, a, b int) int {
	if n == 1 {
		return b
	}
	return f(n-1, b, a+b)
}
