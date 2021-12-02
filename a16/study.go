package a16

import "log"

/*
电话号码组合
2: abc
3: def
4: ghi
5: jkl
6: mno
7: pqrs
8: tuv
9: wxyz
输入: "23"
输出: ["ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"]
*/

func digit2Str(d rune) string {
	if d >= '2' && d <= '6' {
		var c rune
		diff := d - '2'
		c = 'a' + 3 * diff
		return string(c) + string(c + 1) + string(c + 2)
	}
	if d == '7' {
		return "pqrs"
	}
	if d == '8' {
		return "tuv"
	}
	if d == '9' {
		return "wxyz"
	}
	return ""
}

func f(digits string) []string {
	var ret []string
	for _, item := range digits {
		str := digit2Str(item)
		log.Println(str)
		if len(ret) == 0 {
			for _, si := range str {
				ret = append(ret, string(si))
			}
		} else {
			lend := len(str)
			leno := len(ret)
			for i := 0; i < lend - 1; i++ {
				ret = append(ret, ret[0 : leno]...)
			}
			for j := 0; j < lend; j++ {
				for k := 0; k < leno; k++ {
					ret[k + j * leno] = ret[k + j * leno] + string(str[j])
				}
			}
		}
	}
	return ret
}
