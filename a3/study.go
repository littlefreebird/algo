package a3

import "fmt"

/*
找出一个字符串中不含有重复字符的最长子串
输入：ababcdeffa
输出：abcdef
*/

func exist(c byte, set [256]bool) bool {
	return set[int(c)]
}

func f(s string) {
	set := [256]bool{false}
	left := 0
	right := 0
	i := 0
	j := 0
	set[int(s[i])] = true
	stop := 0
	max := 1
	for {
		if j == len(s) - 1  || i == len(s) - 1 {
			fmt.Printf("%v\n", s[left : right + 1])
			break
		}
		if exist(s[j + 1], set) == false {
			j = j + 1
			set[int(s[j])] = true
		} else {
			stop = j + 1
			if j - i + 1 > max {
				left = i
				right = j
				max = j - i + 1
			}
			idx := i
			for idx = i; idx <= stop; idx++ {
				if s[idx] == s[stop] {
					break
				}
			}
			i = idx + 1
			if i > j {
				j = i
			}
			set = [256]bool{false}
			set[int(s[i])] = true
		}
	}
}