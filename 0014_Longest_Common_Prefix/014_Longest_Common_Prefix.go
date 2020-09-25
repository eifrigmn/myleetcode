package _0014

import "strings"

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	if len(strs) == 1 {
		return strs[0]
	}
	var trim = strs[0]
	var s = strs[0]
	var flag = len(s)
	for i := 1; i < len(strs); i++ {
		if len(strs[i]) < flag {
			flag = len(strs[i])
		}
		var tr string
		for j := 0; j < flag; j++ {
			if s[j] != strs[i][j] {
				break
			} else {
				tr = tr + string(s[j])
			}
		}
		if len(tr) < len(trim) {
			trim = tr
		}
	}
	return trim
}

// 最优解
func longestCommonPrefix1(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	prefix := strs[0]
	for _, code := range strs {
		for strings.Index(string(code), prefix) != 0 {
			if len(prefix) == 0 {
				return ""
			}
			prefix = prefix[0 : len(prefix)-1]
		}
	}
	return string(prefix)
}
