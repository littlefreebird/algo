package a54

//找出数组中出现次数超过一般的数字

func f(a []int) int {
	if len(a) == 0 {
		return -1
	}
	if len(a) == 1 {
		return a[0]
	}
	ret := a[0]
	cnt := 1
	for _, item := range a {
		if cnt == 0 {
			cnt = 1
			ret = item
			continue
		}
		if item == ret {
			cnt++
		} else {
			cnt--
		}
	}
	cnt = 0
	for _, item := range a {
		if item == ret {
			cnt++
		}
		if cnt > len(a)/2 {
			return ret
		}
	}
	return -1
}
