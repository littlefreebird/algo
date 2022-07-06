package a35

//给定正整数数组，从位置0开始，最少需要几步到达数组最后一个数字，每一步的大小不能超过当前数字的大小

func f(a []int) int {
	l := len(a)
	d := make([]int, l)
	d[l-1] = 0
	if l == 1 {
		return 0
	}
	for i := l - 2; i >= 0; i-- {
		if a[i] == 0 {
			d[i] = -1
			continue
		}
		step := 1
		min := -1
		for ; step <= a[i]; step++ {
			if l-1-i >= step {
				next := i + step
				if d[next] == -1 {
					min = -1
				} else {
					if min == -1 || 1+d[next] <= min {
						min = 1 + d[next]
					}
				}
			}
		}
		d[i] = min
	}
	return d[0]
}
