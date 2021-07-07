package _0009

func isPalindrome(x int) bool {
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}
	var revertedNumber int
	for x > revertedNumber {
		revertedNumber = revertedNumber*10 + x%10
		x /= 10
	}
	return x == revertedNumber || x == revertedNumber/10
}

// 最优解
func isPalindrome1(x int) bool {
	if x < 0 {
		return false
	} else {
		r := 0
		t := x
		for x > 0 {
			r = r*10 + x%10
			x /= 10
		}
		if r == t {
			return true
		} else {
			return false
		}
	}
}
