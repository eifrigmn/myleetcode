package _09

func longestCommonPrefix(strs []string) string {
	var result string
	if len(strs) == 0 {
		return ""
	}
	for idx := 0;idx <len(strs[0]);idx++{
		for i:=1;i<len(strs);i++{
			if idx >= len(strs[i]) {
				return result
			}

			if strs[i][idx] == strs[0][idx] {
				continue
			} else {
				return result
			}
		}
		result += string(strs[0][idx])
	}
	return result
}
