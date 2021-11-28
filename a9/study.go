package a9

/*判断一个数字是否为回文*/

func f(d int) bool {
	if d < 0 {
		return false
	}
	count := 0
	rest := d
	m := 1
	for {
		count++
		if rest / 10 == 0 {
			break
		}
		m *= 10
		rest = rest / 10
	}
	rest = d
	for i := 0; i < count / 2; i++ {
		if rest / m != rest % m {
			return false
		}
		high := rest / m * m
		low := rest % m
		rest = rest - high - low
		m /= 10
	}
	return true
}