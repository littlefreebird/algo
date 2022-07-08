package a41

//计算一个数字的n次幂

func f(base float64, n int) float64 {
	if n == 1 {
		return base
	}
	if n%2 == 0 {
		x := f(base, n/2)
		return x * x
	} else {
		x := f(base, n/2)
		return x * x * base
	}
}
