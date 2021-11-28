package a13


/*查找字符串数组中最长公共前缀*/

func f(arrStr []string) string {
	length := len(arrStr)
	same := ""
	idx := 0
	for {
		for i := 0; i < length; i++ {
			if idx >= len(arrStr[i]) || arrStr[i][idx] != arrStr[0][idx] {
				return same
			}
		}
		same += arrStr[0][idx:idx+1]
		idx++
	}
}