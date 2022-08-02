package test1

//去除字符串里面的空格

func f(si string) string {
	s := []rune(si)
	r := ' '
	count := 0
	idx := 0
	for {
		if idx > len(s)-1 {
			break
		}
		for {
			if idx >= len(s) {
				break
			}
			if s[idx] != r {
				break
			}
			count++
			idx++
		}
		if count == 0 {
			idx++
		}
		if count > 0 {
			k := 0
			for {
				if idx >= len(s) || s[idx] == r {
					break
				}
				s[idx-count] = s[idx]
				s[idx] = r
				idx++
				k++
			}
			if k > count {
				count = k - count
			} else {
				count = 0
			}
		}
	}
	return string(s)
}
