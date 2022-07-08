package a39

import "fmt"

//将一个n*n的矩阵原地旋转90度

func f(m [][]int, n int) {
	fmt.Println(m)
	idx := 0
	l := n
	for {
		if l <= 1 {
			break
		}
		var tmp []int
		for i := 0; i < 5; i++ {
			for j := 0; j < l; j++ {
				switch i {
				case 0:
					tmp = append(tmp, m[idx][j+idx])
				case 1:
					m[idx][j+idx] = m[n-1-j][idx]
				case 2:
					m[n-i-j][idx] = m[n-idx][n-1-idx-j]
				case 3:
					m[n-idx][n-1-idx-j] = m[n-1-idx][idx+j]
				case 4:
					m[n-1-idx][idx+j] = tmp[j]
				}
			}
		}
		l -= 2
		idx++
	}
	fmt.Println(m)
}