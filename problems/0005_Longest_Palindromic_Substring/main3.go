package _0005

func _longestPalindrome(s string) string {
	var result string
	for i := 0; i < len(s); i++ {
		s1 := getLongestPalindrome(s, i, i)
		s2 := getLongestPalindrome(s, i, i+1)
		result = getMaxSizeStr(result, s1)
		result = getMaxSizeStr(result, s2)
	}
	return result
}

func getLongestPalindrome(s string, l, r int) string {
	for l >= 0 && r < len(s) {
		if s[l] == s[r] {
			l--
			r++
		} else {
			break
		}
	}
	return s[l+1 : r]
}

func getMaxSizeStr(s1, s2 string) string {
	if len(s1) > len(s2) {
		return s1
	}
	return s2
}
