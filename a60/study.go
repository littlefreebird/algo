package a60

//统计一个数字在排序数组中出现的次数

func f(a []int, b int) int {
	return f1(a, 0, len(a)-1, b)
}

func f1(a []int, l int, r int, b int) int {
	if l > r {
		return 0
	}
	mid := (l + r) / 2
	if a[mid] == b {
		return 1 + f1(a, l, mid-1, b) + f1(a, mid+1, r, b)
	} else if a[mid] < b {
		return f1(a, mid+1, r, b)
	} else {
		return f1(a, l, mid-1, b)
	}
}
