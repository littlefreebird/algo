package a43

import "fmt"

//给定一个m*n矩阵，螺旋遍历此矩阵

func f(matrix [][]int, m int, n int) []int {
	ret := make([]int, m*n)
	idx := 0
	round := 0
	for {
		for i := round; i < n-round; i++ {
			ret[idx] = matrix[round][i]
			idx++
		}
		if idx >= m*n {
			break
		}
		for j := round + 1; j < m-round; j++ {
			ret[idx] = matrix[j][n-1-round]
			idx++
		}
		if idx >= m*n {
			break
		}
		for k := n - 2 - round; k > round; k-- {
			ret[idx] = matrix[m-1-round][k]
			idx++
		}
		if idx >= m*n {
			break
		}
		for l := m - 1 - round; l > round; l-- {
			ret[idx] = matrix[l][round]
			idx++
		}
		if idx >= m*n {
			break
		}
		round++
	}
	for _, item := range matrix {
		fmt.Println(item)
	}
	fmt.Println(ret)
	return ret
}
