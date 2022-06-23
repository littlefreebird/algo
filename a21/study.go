package a21

func f(n int) []string {
	if n < 1 || n > 8 {
		return []string{}
	}
	if n == 1 {
		return []string{"()"}
	} else {
		var ret []string
		lastRet := f(n-1)
		for idx, item := range lastRet {
			if idx == len(lastRet) - 1 {
				ret = append(ret, "("+item+")")
				ret = append(ret, "()"+item)
			} else {
				ret = append(ret, "("+item+")")
				ret = append(ret, "()"+item)
				ret = append(ret, item+"()")
			}
		}
		return ret
	}
}