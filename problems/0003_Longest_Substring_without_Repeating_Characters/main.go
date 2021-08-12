package _0003

func lengthOfLongestSubstring(s string) int {
	if len(s) < 2 {
		return len(s)
	}
	var start, end, ans int
	var mp = make(map[uint8]int, len(s))
	for i := 0; end < len(s); i++ {
		_, ok := mp[s[i]]
		if !ok {
			mp[s[i]] = i
		} else {
			start = max(mp[s[i]]+1, start)
			mp[s[i]] = i
		}
		ans = max(end-start+1, ans)
		end++
	}
	return ans
}

func max(num1, num2 int) int {
	if num1 >= num2 {
		return num1
	}
	return num2
}
