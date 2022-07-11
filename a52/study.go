package a52

//给定一个非递减数组的旋转数组，输出数组的最小值，此数组中元素值都大于0

func f(a []int) int {
	if len(a) == 0 {
		return -1
	}
	if a[0] <= a[len(a)-1] {
		return a[0]
	}
	l := 0
	r := len(a) - 1
	for {
		if r-l == 1 {
			return a[r]
		}
		mid := (l + r) / 2
		if a[mid] >= a[l] {
			l = mid
			continue
		}
		if a[mid] <= a[r] {
			r = mid
			continue
		}
	}
}
