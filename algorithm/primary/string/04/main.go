package _04

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	var mps, mpt = make(map[int32]int, 26), make(map[int32]int,26)
	for _, vs := range s {
		mps[vs-'a'] += 1
	}

	for _, vt := range t {
		mpt[vt-'a'] += 1
	}

	for s, tms := range mps{
		if tmst, ok := mpt[s]; ok {
			if tms != tmst {
				return false
			} else {
				continue
			}
		} else {
			return false
		}
	}
	return true
}

func isAnagram1(s string, t string) bool {
	var i int

	var arr [26]int

	if len(s)!=len(t){
		return false
	}

	for i=0;i<len(s);i++{
		arr[s[i]-'a']++
		arr[t[i]-'a']--
	}

	for i=0;i<len(arr);i++{
		if arr[i]!=0{
			return false
		}
	}

	return true
}
