package _05

func isValid(s string) bool {
	var stack []uint8
	var mp = map[uint8]uint8{
		')': '(',
		']': '[',
		'}': '{',
	}
	for i := 0; i < len(s); i++ {
		char, ok := mp[s[i]]
		if len(stack) > 0 && ok {
			top := stack[len(stack)-1]
			if top == char {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		} else {
			stack = append(stack, s[i])
		}
	}

	if len(stack) == 0 {
		return true
	}
	return false
}
