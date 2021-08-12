package _0003

func lengthOfLongestSubstring1(s string) int {
	var result, head, tail int
	if len(s) < 2 {
		return len(s)
	}
	var mp = make(map[uint8]int, len(s))
	for tail < len(s) {
		if idx, ok := mp[s[tail]]; !ok {
			mp[s[tail]] = tail
			tail++
		} else {
			mp = make(map[uint8]int, len(s))
			if tail-head > result {
				result = tail - head
			}
			head = idx + 1
			mp[s[head]] = head
			tail = head + 1
		}
	}
	if tail-head > result {
		result = tail - head
	}
	return result
}
