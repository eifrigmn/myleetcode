package _0005

func longestPalindrome(s string) string {
	for i:=0;i<len(s);i++{
		step := len(s)-i
		for j:=0;j+step<=len(s);j++{
			if isPalindrome(s[j:j+step]) {
				return s[j:j+step]
			}
		}
	}
	return ""

}

func isPalindrome(s string) bool {
	for len(s) > 1 {
		if s[0] != s[len(s)-1] {
			return false
		}
		s = s[1:len(s)-1]
	}
	return true
}

// 0ms
func longestPalindrome1(s string) string {
	length := len(s)

	if length < 2 {
		return s
	}

	start, maxLen := 0, 1
	for i := 0; i < length; {
		if length-i <= maxLen/2 {
			break
		}
		b, e := i, i
		for e < length-1 && s[e+1] == s[e] {
			e++
		}
		i = e + 1
		for e < length-1 && b > 0 && s[e+1] == s[b-1] {
			e++
			b--
		}
		newLen := e + 1 - b
		if newLen > maxLen {
			start = b
			maxLen = newLen
		}
	}
	return s[start : start+maxLen]
}
