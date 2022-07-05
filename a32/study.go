package a32

import "fmt"

/*给定一个正整数n，输出外观数列的第n项
1. 1
2. 11 表示上一项是1个1
3. 21 表示上一项是2个1
4. 1211 表示上一项是1个2，2个1
5. 111221 表示上一项是1个1，1个2，2个1*/

func f(n int) string {
	if n == 1 {
		return "1"
	}
	var ret string
	pre := f(n - 1)
	start := 0
	end := 0
	v := int32(pre[start])
	for i, item := range pre {
		if item == v {
			end = i
		} else {
			ret = ret + fmt.Sprintf("%d%c", end-start+1, v)
			start = i
			end = i
			v = item
		}
	}
	if end == len(pre)-1 {
		ret = ret + fmt.Sprintf("%d%c", end-start+1, v)
	}
	return ret
}
