package a57

import "fmt"

//有一个长度为n的数组，数组中元素值从0到n-1，找出数组中一个重复的数字

func f(a []int) {
	for i, _ := range a {
		if i == a[i] {
			continue
		}
		if a[i] == a[a[i]] {
			fmt.Println(a[i])
			return
		}
		tmp := a[i]
		a[i] = a[tmp]
		a[tmp] = tmp
	}
}
