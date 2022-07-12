package a59

import "fmt"

//输入一个数组，求此数组中逆序对的总数

func f(a []int) {
	if len(a) <= 1 {
		fmt.Println(0)
	}
	var count int
	countReverse(a, 0, len(a)-1, &count)
	fmt.Println(count)
}

func countReverse(a []int, l int, r int, count *int) {
	if l < r {
		mid := (l + r) / 2
		fmt.Println(l, r, mid)
		countReverse(a, l, mid, count)
		countReverse(a, mid+1, r, count)
		merge(a, l, r, mid, count)
	}
}

func merge(a []int, l int, r int, mid int, count *int) {
	i := l
	j := mid + 1
	tmp := make([]int, r-l+1)
	idx := 0
	for {
		if i > mid || j > r {
			break
		}
		if a[i] <= a[j] {
			tmp[idx] = a[i]
			idx++
			i++
		} else {
			tmp[idx] = a[j]
			idx++
			j++
			*count += mid - i + 1
		}
	}
	for {
		if i > mid {
			break
		}
		tmp[idx] = a[i]
		idx++
		i++
	}
	for {
		if j > r {
			break
		}
		tmp[idx] = a[j]
		idx++
		j++
	}
}
