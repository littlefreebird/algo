package a38

import (
	"fmt"
	"sort"
)

//给定一个有重复数字的数组，列出数组的所有排列且不可重复

func f(a []int) {
	sort.Ints(a)
	used := make([]int, len(a))
	one := make([]int, 0, len(a))
	var ret [][]int
	f1(a, used, one, &ret)
	fmt.Println(ret)
}

func f1(a []int, used []int, one []int, ret *[][]int) {
	if len(one) == len(a) {
		dst := make([]int, len(a))
		copy(dst, one)
		*ret = append(*ret, dst)
		return
	}
	for i, _ := range a {
		if used[i] == 0 {
			if i > 0 && used[i-1] == 0 && a[i-1] == a[i] {
				continue
			}
			one = append(one, a[i])
			used[i] = 1
			f1(a, used, one, ret)
			used[i] = 0
			one = one[0 : len(one)-1]
		}
	}
}
