package _0005

// 按照字符串长度递减
// 每次截取指定长度字符串的子字符串
// 判断该字符串若符合回文字符串定义，则直接输出
func longestPalindrome2(s string) string {
	if len(s) < 2 {
		return s
	}
	if isPalindrome2(s) {
		return s
	}
	var psize = len(s) - 1
	for psize > 0 {
		for i := len(s) - 1; i >= psize-1; i-- {
			lft := i - psize + 1
			tmp := cutString(s, lft, i)
			if isPalindrome2(tmp) {
				return tmp
			}
		}
		psize--
	}

	return string(s[0])
}

// 判定一个回文字符串
func isPalindrome2(s string) bool {
	if len(s) < 2 {
		return false
	}

	for len(s) > 1 {
		if s[0] != s[len(s)-1] {
			return false
		}

		s = s[1 : len(s)-1]
	}
	return true
}

func cutString(s string, lft, rgt int) string {
	if rgt == len(s)-1 {
		return s[lft:]
	}
	return s[lft : rgt+1]
}
