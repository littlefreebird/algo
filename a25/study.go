package a25

//在一个字符串中找出第一次出现的位置

func next(sub string) []int {
	var ret []int
	j := 0
	for i := 0; i < len(sub); i++ {
		if i == 0 || i == 1 {
			ret = append(ret, 0)
			continue
		}
		if sub[i-1] == sub[j] {
			ret = append(ret, j + 1)
			j++
		} else {
			for {
				j = ret[j]
				if j == 0 {
					ret = append(ret, 0)
					break
				} else {
					if sub[j] == sub[i-1] {
						ret = append(ret, j + 1)
						j++
						break
					}
				}
			}
		}
	}
	return ret
}

func kmp(s string, sub string) int {
	na := next(sub)
	i := 0
	j := 0
	if len(s) < 1 || len(sub) < 1 {
		return -1
	}
	for {
		if s[i] == sub[j] {
			i++
			j++
		} else {
			j = na[j]
		}
		if j >= len(sub) && i <= len(s) {
			return i-len(sub)
		}
		if j >= len(sub) && i > len(s) {
			return -1
		}
	}
}