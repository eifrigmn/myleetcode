package _0028

func strStr(haystack string, needle string) int {
	if needle == "" {
		return 0
	}
	if len(haystack) <len(needle) {
		return -1
	}
	for i:=0;i<len(haystack);i++ {
		if len(haystack[i:]) <len(needle) {
			return -1
		}
		if haystack[i:i+len(needle)] == needle {
			return i
		}
	}
	return -1
}

// 参考
func strStr1(haystack string, needle string) int {
	lh := len(haystack)
	ln := len(needle)
	if lh < ln {
		return -1
	}
	if ln == 0 {
		return 0
	}

	flag := 0
	for flag <= lh-ln {
		if haystack[flag:flag+ln] == needle {
			return flag
		} else {
			flag += ln
			if flag >= lh {
				return -1
			}
			has := false
			for i := ln - 1; i >= 0; i-- {
				if needle[i:i+1] == haystack[flag:flag+1] {
					flag -= i
					has = true
					break
				}
			}
			if !has {
				flag++
			}
		}
	}
	return -1
}