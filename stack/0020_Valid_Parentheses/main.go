package stack

func isValid(s string) bool {
	pre := make([]rune, len(s))
	i := -1

	for _, c := range s {

		if i > -1 {
			if c == ']' && pre[i] == '[' || c == ')' && pre[i] == '(' || c == '}' && pre[i] == '{' {
				i--
				continue
			}
		}

		i++
		pre[i] = c
	}

	return i == -1

}

// 参考
func isValid1(s string) bool {
	if len(s) == 0 {
		return true
	}
	stack := []string{}
	counterPart := map[string]string{
		"}": "{",
		"]": "[",
		")": "(",
	}

	for i := range s {
		current := string(s[i])

		if len(stack) != 0 && counterPart[current] != "" {
			top := stack[len(stack)-1]
			if top == counterPart[current] {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		} else {
			stack = append(stack, current)
		}
	}

	if len(stack) == 0 {
		return true
	}

	return false
}
