package a48

import "fmt"

//给定一个正整数n，生成一个包含1到n^2所有元素，且元素按顺时针顺序螺旋排列的n*n正方形矩阵

func f(n int) {
	m := make([][]int, n)
	for idx, _ := range m {
		m[idx] = make([]int, n)
	}
	idx := 1
	for i := 0; i < n/2; i++ {
		step := n - 2*i - 1
		for j := 0; j < step; j++ {
			m[i][j+i] = idx
			idx++
		}
		for j := 0; j < step; j++ {
			m[j+i][n-1-i] = idx
			idx++
		}
		for j := 0; j < step; j++ {
			m[n-1-i][n-1-i-j] = idx
			idx++
		}
		for j := 0; j < step; j++ {
			m[n-1-i-j][i] = idx
			idx++
		}
	}
	if n%2 == 1 {
		m[n/2][n/2] = idx
	}
	for i := 0; i < n; i++ {
		fmt.Println(m[i])
	}
}
