package a37

import "fmt"

//给定一个没有重复数字的数组，计算数组的全排列

func f(a []int) {
	used := make([]int, len(a))
	var ret [][]int
	one := make([]int, 0, len(a))
	f1(a, used, one, &ret)
	fmt.Println(ret)
}

func f1(a []int, used []int, one []int, ret *[][]int) {
	if len(one) == len(a) {
		dst := make([]int, len(one))
		copy(dst, one)
		*ret = append(*ret, dst)
		return
	}
	for i := 0; i < len(a); i++ {
		if used[i] == 0 {
			one = append(one, a[i])
			used[i] = 1
			f1(a, used, one, ret)
			used[i] = 0
			one = one[0 : len(one)-1]
		}
	}
}
