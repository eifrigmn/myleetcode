package _01

// 1	1
// 2	2
// 3	3
// 4	5
// 5	8
// f(n+2) = f(n+1) + f(n)
// f(1) = 1
// f(2) = 2
func climbStairs(n int) int {
	return Fibonacci(n, 1, 1)
}

func Fibonacci(n, a, b int) int {
	if n <= 1 {
		return b
	}
	return Fibonacci(n-1, b, a+b)
}
