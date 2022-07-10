package a50

/*一个机器人位于一个m*n网格的左上角，机器人每次只能往下或往右移动一步，机器人试图到达网格的右下角
请问有多少种路径*/

func f(m int, n int) int {
	if m == 1 && n == 2 {
		return 1
	}
	if m == 2 && n == 1 {
		return 1
	}
	if m == 1 && n == 1 {
		return 0
	}
	if m < 1 || n < 1 {
		return 0
	}
	return f(m, n-1) + f(m-1, n)
}
