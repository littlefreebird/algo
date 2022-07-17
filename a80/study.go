package a80

//我们可以用2*1的矩阵去覆盖更大的矩阵，请问用n个2*1的矩阵去覆盖一个2*n的大矩阵，有多少种方法

func f(n int) int {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	return f(n-1) + f(n-2)
}
