package _04

func romanToInt(s string) int {
	var mp = map[uint8]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	var result int
	for len(s) > 0 {
		if len(s) > 1 && (s[:2] == "IV" || s[:2] == "IX" || s[:2] == "XL" || s[:2] == "XC" || s[:2] == "CD" || s[:2] == "CM") {
			result = result + mp[s[1]] - mp[s[0]]
			s = s[2:]
		} else {
			result = result + mp[s[0]]
			s = s[1:]
		}
	}
	return result
}
