package a83

/*给定一个字符串
1. 如果字符串为123.4K，返回123.4*1024的的整数值
2. 如果字符串为23M，返回23*1024*1024的整数值
3. 如果字符串为98G，返回98*1024*1024*1024的整数值
4. 其他情况返回-1*/

import "fmt"

func f(s string) int {
	diff := 0 - '0'
	length := len(s)
	sum1 := 0
	var sum2 float64 = 0
	dot := 0
	count := 10
	mul := 1
	con := true
	for idx, item := range s {
		if dot > 1 {
			return -1
		}
		if idx == 0 && isDigit(item) == false {
			return -1
		}
		if idx > 0 {
			if idx == length-1 {
				if item == 'K' {
					mul = 1024
					con = false
				} else if item == 'M' {
					mul = 1024 * 1024
					con = false
				} else if item == 'G' {
					mul = 1024 * 1024 * 1024
					con = false
				} else if isDigit(item) {
					con = true
				} else {
					return -1
				}
			}
			if item == '.' {
				dot = 1
				continue
			}
			if idx != length-1 && isDigit(item) == false {
				return -1
			}
		}
		if con {
			if dot == 0 {
				sum1 = sum1*10 + int(item+diff)
			} else {
				sum2 = sum2 + float64(item+diff)/float64(count)
				count = count * 10
			}
		}
	}
	ret := int((float64(sum1) + sum2) * float64(mul))
	fmt.Println(s, ret)
	return ret
}

func isDigit(i int32) bool {
	if i >= '0' && i <= '9' {
		return true
	}
	return false
}
