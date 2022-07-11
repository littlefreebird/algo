package a51

//一个二维数组，横向众向都单调递增，判断二维数组中是否存在一个给定的数

func f(m [][]int, l int, r int, t int, b int, x int) (int, int) {
	for {
		if l > r || t > b {
			return -1, -1
		}
		hm := (l + r) / 2
		if m[t][hm] == x {
			return t, hm
		} else if x > m[t][hm] {
			l = hm + 1
		} else {
			r = hm - 1
		}
		vm := (t + b) / 2
		if m[t][vm] == x {
			return t, vm
		} else if x > m[t][vm] {
			t = vm + 1
		} else {
			b = vm - 1
		}
	}
}
