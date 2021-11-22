package a6

/*
给定一个字符串s1和一个字符串s2，找出s1中第一个这样的字符串s3：s3是s2中所有字符的一种排列
比如s1=abcdef，s2=cba，那么s3=abc
比如s1=xyzlmn，s2=nm，那么s3=mn
*/


func f(s1 string, s2 string) string {
	len := 0
	var total [256]rune
	for _, item2 := range s2 {
		total[item2] = total[item2] + 1
		len = len + 1
	}
	var tmpTotal = total
	realLen := 0
	start := 0
	end := 0
	for idx, item1 := range s1 {
		if tmpTotal[item1] > 0 {
			tmpTotal[item1] = tmpTotal[item1] - 1
			end = idx
			realLen = realLen + 1
			if realLen == len {
				return s1[start:end+1]
			}
		} else {
			start = idx + 1
		}
	}
	return ""
}