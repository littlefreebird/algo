package a31

/*判断一个9*9的数独是否有效
1.数字1-9在每一行只能出现一次
2.数字1-9在每一列只能出现一次
3.数字1-9在99个3*3的宫内只能出现一次*/

func f(a [9][9]byte) bool {
	var row [9][9]int
	var col [9][9]int
	var block [9][9]int
	for i, _ := range a {
		for j, _ := range a[i] {
			if a[i][j] == '.' {
				continue
			}
			d := a[i][j] - 1
			row[i][d] = row[i][d] + 1
			col[j][d] = col[j][d] + 1
			block[i/3*3+j/3][d] = block[i/3*3+j/3][d] + 1
			if row[i][d] > 1 || col[j][d] > 1 || block[i/3*3+j/3][d] > 1 {
				return false
			}
		}
	}
	return true
}
