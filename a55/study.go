package a55

import (
	"fmt"
	"strconv"
)

/*给定一个正整数数组，把数组中所有数字拼接成一个数，找出值最小的拼接数字
比如{3, 321, 32}，最小的拼接数字就是321323*/

func f(a []int) string {
	fmt.Println(a)
	if len(a) == 0 {
		return ""
	}
	if len(a) == 1 {
		return strconv.Itoa(a[0])
	}
	quickSort(a, 0, len(a)-1)
	ret := ""
	for _, item := range a {
		ret += strconv.Itoa(item)
	}
	return ret
}

func quickSort(a []int, l, r int) {
	if l >= r {
		return
	}
	x := a[l]
	i := l
	j := r
	for {
		if i >= j {
			break
		}
		for {
			if j > i && !cmp(a[j], x) {
				j--
				continue
			}
			break
		}
		if j > i {
			a[i] = a[j]
			i++
		}
		for {
			if i < j && cmp(a[i], x) {
				i++
				continue
			}
			break
		}
		if j > i {
			a[j] = a[i]
			j--
		}
	}
	a[i] = x
	quickSort(a, l, i-1)
	quickSort(a, i+1, r)
}

func cmp(a, b int) bool {
	i := 10
	for {
		if a/i == 0 {
			break
		}
		i = i * 10
	}
	j := 10
	for {
		if b/j == 0 {
			break
		}
		j = j * 10
	}
	if a*j+b <= b*i+a {
		return true
	}
	return false
}
