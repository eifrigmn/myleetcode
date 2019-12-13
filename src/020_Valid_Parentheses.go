/*
题目：有效的括号
	给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效。
	有效字符串需满足：
		左括号必须用相同类型的右括号闭合。
		左括号必须以正确的顺序闭合。
		注意空字符串可被认为是有效字符串。

	示例 1:
		输入: "()"
		输出: true
	示例 2:
		输入: "()[]{}"
		输出: true
	示例 3:
		输入: "(]"
		输出: false
	示例 4:
		输入: "([)]"
		输出: false
	示例 5:
		输入: "{[]}"
		输出: true
来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/valid-parentheses
*/
package src

func isValid(s string) bool {
	bracketsMap := map[uint8]uint8{'{': '}', '[': ']', '(': ')'}
	if s == "" {
		return true
	}
	if len(s)/2 != 0 {
		return false
	}

	var stack []uint8
	for i := 0; i < len(s); i++ {
		if len(stack) > 0 {
			if c, ok := bracketsMap[stack[len(stack)-1]]; ok && c == s[i] {
				stack = stack[0 : len(stack)-2]
				continue
			}
		}
		stack = append(stack, s[i])
	}
	return len(stack) == 0
}

// 最优解
func isValid1(s string) bool {
	pre := make([]rune, len(s))
	i := -1

	for _, c := range s {

		if i > - 1 {
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
