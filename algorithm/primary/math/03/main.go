package _03

func isPowerOfThree(n int) bool {
	if n == 0 {
		return false
	}

	if n == 1 {
		return true
	}
	return isPowerOf(n, 3)
}

func isPowerOf(n, tail int) bool {
	if n < tail {
		return false
	} else if n == tail {
		return true
	} else {
		return isPowerOf(n, tail*3)
	}
}
