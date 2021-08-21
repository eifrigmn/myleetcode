package _0008

import "strings"

func myAtoi(s string) int {
	var result int64
	s = strings.TrimSpace(s)
	for len(s) > 1 {
		if s[len(s)-1] <= '9' && s[len(s)-1] >= '0' {
			result = result*10 + int64((s[len(s)-1] - '0'))
		}
		s = s[:len(s)-1]
	}
}
