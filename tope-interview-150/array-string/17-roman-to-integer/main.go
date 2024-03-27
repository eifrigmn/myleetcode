package main

func romanToInt(s string) int {
	if len(s) == 1 {
		return getIngeter(s[0])
	}
	var result int
	for i := 0; i < len(s)-1; i++ {
		if getIngeter(s[i]) < getIngeter(s[i+1]) {
			result = result - getIngeter(s[i])
		} else {
			result = result + getIngeter(s[i])
		}
	}
	result += getIngeter(s[len(s)-1])
	return result
}

func getIngeter(roman byte) int {
	switch roman {
	case byte('I'):
		return 1
	case byte('V'):
		return 5
	case byte('X'):
		return 10
	case byte('L'):
		return 50
	case byte('C'):
		return 100
	case byte('D'):
		return 500
	case byte('M'):
		return 1000
	}
	return 0
}
