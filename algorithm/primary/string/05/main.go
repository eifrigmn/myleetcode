package _05

import (
	"strings"
)

func isPalindrome(s string) bool {
	s = strings.ToLower(s)
	var lft, rgt = 0, len(s)-1
	for lft <rgt {
		if !(s[lft] >= '0' && s[lft] <= '9') && !(s[lft] >= 'a' && s[lft] <= 'z') {
			lft++
			continue
		}

		if !(s[rgt] >= '0' && s[rgt] <= '9') && !(s[rgt] >= 'a' && s[rgt] <= 'z') {
			rgt--
			continue
		}

		if s[lft] != s[rgt] {
			return false
		}

		lft++
		rgt--
	}
	return true
}