package _0003

func _lengthOfLongestSubstring(s string) int {
	var tmp = make(map[uint8]int)
	var length int
	for i, j := 0, 0; j < len(s); j++ {
		lastIdx, ok := tmp[s[j]]
		if ok {
			if lastIdx+1 > i {
				i = lastIdx + 1
			}
		}
		if j-i+1 > length {
			length = j - i + 1
		}
		tmp[s[j]] = j
	}
	return length
}
