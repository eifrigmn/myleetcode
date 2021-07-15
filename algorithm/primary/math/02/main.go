package _02

func countPrimes(n int) int {
	if n == 0 || n == 1 {
		return 0
	}
	var cnt = n - 2
	var primes = make([]bool, n+1)
	for i := 2; i < n; i++ {
		for j := 2; i*j < n; j++ {
			if primes[i*j] == false {
				primes[i*j] = true
				cnt--
			}
		}
	}
	return cnt
}
