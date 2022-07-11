package a53

import "fmt"

//给定一个数组，将数组中奇数移动到数组的前半部分，将偶数移动到数组的后半部分，且奇数之间的相对位置不变偶数之间的相对位置不变

func f(a []int) {
	fmt.Println(a)
	j := 0
	for i := 0; i < len(a); i++ {
		if a[i]%2 == 1 {
			continue
		} else {
			j = i
			for {
				j = j + 1
				if j >= len(a) {
					break
				}
				if a[j]%2 == 1 {
					tmp := a[i]
					a[i] = a[j]
					a[j] = tmp
					break
				}
			}
		}
	}
	fmt.Println(a)
}
