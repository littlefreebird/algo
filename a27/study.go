package a27

import "fmt"

//给定一个排列，输出此排列的下一个相邻排列

func f(a []int) []int {
	ret := make([]int, len(a))
	min := 2 << 16
	max := 0
	minIdx := len(a) - 1
	maxIdx := minIdx
	for i := len(a) - 1; i >= 0; i-- {
		if a[i] < max {
			fmt.Println(i, minIdx)
			ret[i] = min
			ii := i + 1
			for j := len(a) - 1; j > minIdx; j++ {
				ret[ii] = a[j]
				ii = ii + 1
			}
			ret[ii] = a[i]
			for k := minIdx - 1; k > i; k-- {
				ii = ii + 1
				ret[ii] = a[k]
			}
			for l := 0; l < i; l++ {
				ret[l] = a[l]
			}
			break
		}
		if a[i] > max {
			max = a[i]
			maxIdx = i
		}
		if a[i] < min {
			min = a[i]
			minIdx = i
		}

	}
	if maxIdx == 0 {
		for i := len(a); i > 0; i-- {
			ret[i-1] = i
		}
	}
	fmt.Println(a)
	fmt.Println(ret)
	return ret
}
