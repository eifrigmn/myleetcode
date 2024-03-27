package main

func reverseWords(s string) string {
	var result string
	for len(s) > 1 && s[0] == byte(' ') {
		s = s[1:]
	}
	s = " " + s
	for len(s) > 0 && s[len(s)-1] == byte(' ') {
		s = s[:len(s)-1]
	}
	var headIdx, tailIdx = len(s) - 1, len(s) - 1
	for headIdx >= 0 {
		if s[headIdx] == byte(' ') {
			if headIdx != tailIdx {
				result = result + s[headIdx+1:tailIdx+1] + " "
				tailIdx = headIdx
			} else {
				tailIdx--
				headIdx--
			}
		} else {
			headIdx--
		}
	}
	if len(result) > 0 && result[len(result)-1] == byte(' ') {
		result = result[:len(result)-1]
	}
	return result
}
