package main

func lengthOfLastWord(s string) int {
	var lastWordCnt int
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == byte(' ') {
			if lastWordCnt > 0 {
				break
			} else {
				continue
			}
		} else {
			lastWordCnt++
		}
	}
	return lastWordCnt
}
