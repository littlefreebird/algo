package a58

import "fmt"

//给定一个长度为n的数组A，声场一个数组B，且B中元素i的值B[i]=A[0]*...*A[i-1]*A[i+1]*...*A[n-1]，不能使用除法

func f(A []int) {
	B := make([]int, len(A))
	for i, _ := range A {
		if i == 0 {
			B[i] = 1
		} else {
			B[i] = B[i-1] * A[i-1]
		}
	}
	tmp := 1
	for i := len(A) - 2; i >= 0; i-- {
		tmp = tmp * A[i+1]
		B[i] = B[i] * tmp
	}
	fmt.Println(A)
	fmt.Println(B)
}
