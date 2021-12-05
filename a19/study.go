package a19

/*
判断由()[]{}组成的字符串是否一个有效括号
输入: (){}[]
输出: true

输入: (())
输出: true

输入: ({})[(])
输出: false
*/

func legal(c rune) bool {
	if c == '(' || c == ')' || c == '[' || c == ']' || c == '{' || c == '}' {
		return true
	}
	return false
}

func left(c rune) bool {
	if c == '(' || c == '[' || c == '{' {
		return true
	}
	return false
}

func f(s string) bool {
	var stack []rune
	idx := 0
	for _, i := range s {
		if false == legal(i) {
			return false
		}
		if left(i) {
			if len(stack) - 1 < idx {
				stack = append(stack, i)
			} else {
				stack[idx] = i
			}
			idx++
		}
		if i == ')' {
			if idx > 0 && stack[idx - 1] == '(' {
				idx--
			} else {
				return false
			}
		}
		if i == ']' {
			if idx > 0 && stack[idx - 1] == '[' {
				idx--
			} else {
				return false
			}
		}
		if i == '}' {
			if idx > 0 && stack[idx - 1] == '{' {
				idx--
			} else {
				return false
			}
		}
	}
	return true
}
