package main

func longestCommonPrefix(strs []string) string {
	if len(strs) == 1 {
		return strs[0]
	}
	commonPrefix := getLongestCommonPrefix(strs[0], strs[1])
	for i := 2; i < len(strs); i++ {
		commonPrefix = getLongestCommonPrefix(commonPrefix, strs[i])
		if commonPrefix == "" {
			break
		}
	}
	return commonPrefix
}

func getLongestCommonPrefix(str1, str2 string) string {
	n := len(str1)
	if len(str2) < n {
		n = len(str2)
	}
	var prefix string
	for i := 0; i < n; i++ {
		if str1[i] != str2[i] {
			break
		} else {
			prefix = prefix + string(str1[i])
		}
	}
	return prefix
}
