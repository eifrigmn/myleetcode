package _09

func longestCommonPrefix1(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	if len(strs) == 1 {
		return strs[0]
	}
	var prefix = commonPrefix(strs[0], strs[1])
	for i := 2; i < len(strs); i++ {
		if prefix == "" {
			return prefix
		}
		prefix = commonPrefix(prefix, strs[i])
	}
	return prefix
}

func commonPrefix(a, b string) string {
	var i int
	for ; i < len(a); i++ {
		if i >= len(b) {
			return b
		}
		if a[i] != b[i] {
			return a[:i]
		}
	}
	return a
}
