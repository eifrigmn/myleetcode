package _0009

func isPalindrome2(x int) bool {
	if x < 0 {
		return false
	}

	// 既然是回文数字，则把数字x反转一遍，其值应该不变
	t := x
	r := 0
	for x > 0 {
		r = r*10 + x%10
		x = x / 10
	}
	return r == t

}
