package _07

func strStr(haystack string, needle string) int {
	if needle == "" {
		return 0
	}
	for i:=0;i<len(haystack);i++{
		if len(haystack)-i < len(needle) {
			return -1
		}

		if haystack[i:len(needle)+i] == needle {
			return i
		} else {
			continue
		}
	}
	return -1
}