package a56

import "fmt"

//一个数组中，有两个数字都只有一个，其他数字在数组中都存在两个，找出这两个数字

func f(arr []int) {
	var sum int
	for _, item := range arr {
		sum ^= item
	}
	cnt := 0
	for {
		shift := 1 << cnt
		if (sum & shift) != 0 {
			break
		}
		cnt++
	}
	var a int
	var b int
	for _, item := range arr {
		shift := 1 << cnt
		if (item & shift) != 0 {
			a = a ^ item
		} else {
			b = b ^ item
		}
	}
	fmt.Println(a, b)
}
