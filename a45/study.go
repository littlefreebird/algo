package a45

//一个字符串由多个单词组成，单词之间存在多个空格，求最后一个单词的长度

func f(s string) int {
	flag := false
	l := 0
	i := len(s) - 1
	for ; i >= 0; i-- {
		if flag == false && s[i] != ' ' {
			flag = true
			l++
			continue
		}
		if flag {
			if s[i] != ' ' {
				l++
				if i == 0 {
					return l
				}
			} else {
				return l
			}
		}
	}
	return -1
}
