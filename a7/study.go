package a7

import "log"

/*
给定一个字符串和一个数字n，将字符按宽度为n的z字形输出，比如：
abcdefghijk, 3
a   e   i
b d f h j
c   g   k
输出：aeibdfhjcgk
*/

func f(s string, n int) string {
	len := len(s)
	i := 0
	res := ""
	for {
		pos := (2 * n - 2) * i
		if n == 1 {
			pos = i
		}
		if pos >= len{
			break
		}
		res = res + s[pos:pos+1]
		i = i + 1
	}
	log.Println(res)
	for j := 1; j < n - 1; j++ {
		res = res + s[j:j+1]
		i := 1
		for {
			pos := (2 * n - 2) * i
			if pos - j >= len {
				break
			}
			res = res + s[pos-j:pos-j+1]
			if pos + j >= len {
				break
			}
			res = res + s[pos+j:pos+j+1]
			i = i + 1
		}
	}
	log.Println(res)
	if n < 2 {
		return res
	}
	i = 0
	for {
		pos := n - 1 + (n * 2 - 2) * i
		if pos >= len {
			break
		}
		res = res + s[pos:pos+1]
		i = i + 1
	}
	return res
}