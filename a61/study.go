package a61

//实现一个函数用来匹配包含.和*的正则表达式

func f(s string, p string) bool {
	return f1(s, p, 0, 0)
}

func f1(s string, p string, si int, pi int) bool {
	if si == len(s) && pi == len(p) {
		return true
	}
	if si == len(s) && pi == len(p)-2 && p[pi+1] == '*' {
		return true
	}
	if si >= len(s) || pi >= len(p) {
		return false
	}
	if pi < len(p)-1 && p[pi+1] == '*' {
		if p[pi] == s[si] {
			si++
			return f1(s, p, si, pi)
		} else {
			pi += 2
			return f1(s, p, si, pi)
		}
	} else {
		if p[pi] == '.' || p[pi] == s[si] {
			pi++
			si++
			return f1(s, p, si, pi)
		} else {
			return false
		}
	}
}
