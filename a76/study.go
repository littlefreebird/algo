package a76

/*设计一个函数，判断在一个矩阵中是否存在一条包含某字符串中所有字符的路径。路径可以从矩阵中任意一个字符开始，每一步可以在
矩阵中向上、向下、向左、向右移动。如果一条路径经过了一个格子，下次不能再次进入此格子*/

func f(a [][]byte, m int, n int, s string) bool {
	used := make([][]int, m)
	for i := 0; i < m; i++ {
		used[i] = make([]int, n)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if f1(a, used, m, n, i, j, s, 0) == true {
				return true
			}
		}
	}
	return false
}

func f1(a [][]byte, used [][]int, m int, n int, i int, j int, s string, idx int) bool {
	if used[i][j] == 1 || a[i][j] != s[idx] {
		return false
	}
	used[i][j] = 1
	idx++
	if idx == len(s)-1 {
		return true
	}
	r := false
	if i > 0 {
		r = r || f1(a, used, m, n, i-1, j, s, idx)
	}
	if i < m-1 {
		r = r || f1(a, used, m, n, i+1, j, s, idx)
	}
	if j > 0 {
		r = r || f1(a, used, m, n, i, j-1, s, idx)
	}
	if j < n-1 {
		r = r || f1(a, used, m, n, i, j+1, s, idx)
	}
	if r == false {
		idx--
		used[i][j] = 0
	}
	return r
}
