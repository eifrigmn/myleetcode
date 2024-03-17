package _0350

func reverseString(s []byte) {
	var lft, rgt = 0, len(s) - 1
	for lft < rgt {
		s[lft], s[rgt] = s[rgt], s[lft]
		lft++
		rgt--
	}
}
