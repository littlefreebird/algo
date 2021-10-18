package a4
/*
给定一个字符串，找出其中最长回文子串
输入：s = "abcbx"
输出："bcb"

输入：s = "a"
输出："a"

输入：s = "baax"
输出："aa"
*/

func f(s string) string {
	left := 0
	right := 0
	max := 0
	for idx, item := range s {
		i := idx
		j := idx
		for {
			if i > 0 && s[i - 1] == uint8(item)  {
				i = i - 1
			} else {
				break
			}
		}
		for {
			if j < len(s) - 1 && s[j + 1] == uint8(item) {
				j = j + 1
			} else {
				break
			}
		}
		for {
			if i > 0 && j < len(s) - 1 && s[i - 1] == s[j + 1] {
				i = i - 1
				j = j + 1
			} else {
				break
			}
		}
		if j - i + 1 > max {
			max = j - i + 1
			left = i
			right = j
		}
	}
	return s[left : right + 1]
}
