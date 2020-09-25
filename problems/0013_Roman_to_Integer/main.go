package _0013

import "strings"

func romanToInt(s string) int {
	var sum = 0
	var i = 0
	for i < len(s) {
		if (string(s[i]) == "I" || string(s[i]) == "X" || string(s[i]) == "C") && i != len(s)-1 {
			if negativeFlag(string(s[i+1]), string(s[i])) {
				sum = sum - romanFlagToInt(string(s[i]))
				sum = sum + romanFlagToInt(string(s[i+1]))
				i = i + 2
				continue
			}
		}
		sum = sum + romanFlagToInt(string(s[i]))
		i++
	}
	return sum
}

func romanFlagToInt(s string) int {
	switch strings.ToUpper(s) {
	case "I":
		return 1
	case "V":
		return 5
	case "X":
		return 10
	case "L":
		return 50
	case "C":
		return 100
	case "D":
		return 500
	case "M":
		return 1000
	default:
		return 0
	}
}

func negativeFlag(nextChar, char string) bool {
	if char == "I" && (nextChar == "V" || nextChar == "X") {
		return true
	}

	if char == "X" && (nextChar == "L" || nextChar == "C") {
		return true
	}

	if char == "C" && (nextChar == "D" || nextChar == "M") {
		return true
	}
	return false
}

// 最优解
func romanToInt1(s string) int {
	m := map[rune]int{
		'M': 1000,
		'D': 500,
		'C': 100,
		'L': 50,
		'X': 10,
		'V': 5,
		'I': 1,
	}
	res := 0
	rArray := []rune(s)
	pre := 0
	length := len(rArray)

	for i := length - 1; i >= 0; i-- {
		r := m[rArray[i]]
		if r >= pre {
			res += r
		}
		if r < pre {
			res -= r
		}
		pre = r
	}
	return res
}
