package a69

//给定一个整数数组，判断数组是否某二叉树的后续遍历的结果

func f(a []int) bool {
	return f1(a, 0, len(a)-1)
}

func f1(a []int, left int, right int) bool {
	if left == right {
		return true
	}
	t := a[right]
	i := left
	for ; i < right; i++ {
		if a[i] > t {
			break
		}
	}
	for j := i; j < right; j++ {
		if a[j] < t {
			return false
		}
	}
	return f1(a, left, i-1) && f1(a, i, right)
}
