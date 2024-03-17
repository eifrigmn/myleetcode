package main

// you can also use imports, for example:
// import "fmt"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

import (
	"strings"
)

func Solution(S string) int {
	// Implement your solution here
	var result = -1
	strs := strings.Split(S, " ")
	for _, s := range strs {
		l := passwordCnt(s)
		if l > result {
			result = l
		}
	}
	return result
}

func passwordCnt(s string) int {
	var strCnt, numCnt int
	for i := 0; i < len(s); i++ {
		if isValidChar(s[i]) {
			strCnt++
		} else if isValidNum(s[i]) {
			numCnt++
		} else {
			return -1
		}
	}
	if strCnt%2 == 0 && numCnt%2 == 1 {
		return len(s)
	}
	return -1
}

func isValidChar(c byte) bool {
	return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z')
}

func isValidNum(c byte) bool {
	return c >= '0' && c <= '9'
}
