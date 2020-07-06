package _03

func firstUniqChar(s string) int {
	var mp = make(map[uint8]int, len(s))
	for i:=0;i<len(s);i++{
		if _, ok := mp[s[i]]; ok {
			mp[s[i]] += 1
			continue
		}
		mp[s[i]] = 1
	}
	for i:=0;i<len(s); i++{
		if t := mp[s[i]]; t == 1 {
			return i
		}
	}
	return -1
}

func firstUniqChar1(s string) int {
	last := [26]int{}

	for i, v := range s {
		last[v-'a'] = i
	}

	for i, v := range s {
		if last[v-'a'] == i {
			return i
		} else {
			last[v-'a'] = -1
		}
	}
	return -1
}